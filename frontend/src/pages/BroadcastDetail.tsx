import {
  Badge,
  Button,
  Checkbox,
  Datepicker,
  FileInput,
  Label,
  Modal,
  Pagination,
  TabItem,
  Table,
  Tabs,
  Textarea,
  TextInput,
  ToggleSwitch,
  Tooltip,
} from "flowbite-react";
import moment from "moment";
import { useContext, useEffect, useRef, useState, type FC } from "react";
import Chart from "react-google-charts";
import toast from "react-hot-toast";
import { BsCheck2Circle, BsImage } from "react-icons/bs";
import { FaXmark } from "react-icons/fa6";
import { HiMagnifyingGlass } from "react-icons/hi2";
import Moment from "react-moment";
import { useParams } from "react-router-dom";
import Select from "react-select";
import AdminLayout from "../components/layouts/admin";
import MessageTemplateField from "../components/MessageTemplateField";
import ModalProductList from "../components/ModalProductList";
import { LoadingContext } from "../contexts/LoadingContext";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { BroadcastModel } from "../models/broadcast";
import { ConnectionModel } from "../models/connection";
import { ContactModel } from "../models/contact";
import { FileModel } from "../models/file";
import { ProductModel } from "../models/product";
import { TagModel } from "../models/tag";
import { PaginationResponse } from "../objects/pagination";
import {
  addContactBroadcast,
  addContactFromFileBroadcast,
  deleteContactBroadcast,
  getBroadcast,
  sendBroadcast,
  updateBroadcast,
} from "../services/api/broadcastApi";
import { uploadFile } from "../services/api/commonApi";
import { getConnections } from "../services/api/connectionApi";
import { countContactByTag, getContacts } from "../services/api/contactApi";
import { getProducts } from "../services/api/productApi";
import { getContrastColor, money } from "../utils/helper";
import { parseMentions } from "../utils/helper-ui";
import { IoDocumentsOutline } from "react-icons/io5";
import { GoClock } from "react-icons/go";
import { AiOutlineEdit } from "react-icons/ai";

interface BroadcastDetailProps {}
const neverMatchingRegex = /($a)/;
const BroadcastDetail: FC<BroadcastDetailProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const [emojis, setEmojis] = useState([]);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { broadcastId } = useParams();
  const [mounted, setMounted] = useState(false);
  const [broadcast, setBroadcast] = useState<BroadcastModel>();
  const [isEditable, setisEditable] = useState(false);
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const [showModal, setShowModal] = useState(false);
  const [tags, setTags] = useState<TagModel[]>([]);
  const [selectedTags, setSelectedTags] = useState<TagModel[]>([]);
  const [contacts, setContacts] = useState<ContactModel[]>([]);
  const [selectedContacts, setSelectedContacts] = useState<ContactModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [modalProduct, setModalProduct] = useState(false);
  const [products, setProducts] = useState<ProductModel[]>([]);
  const [selectedProducts, setSelectedProducts] = useState<ProductModel[]>([]);
  const [files, setFiles] = useState<FileModel[]>([]);
  const fileRef = useRef<HTMLInputElement>(null);
  const [selectedBroadcastContacts, setSelectedBroadcastContacts] = useState<
    ContactModel[]
  >([]);
  const [scheduledTimeDistance, setScheduledTimeDistance] = useState(0);

  const [countdown, setCountdown] = useState("");

  const countdownToScheduledAt = (scheduledAt: string) => {
    const targetDate = new Date(scheduledAt).getTime();
    const now = new Date().getTime();
    const distance = targetDate - now;
    setScheduledTimeDistance(distance);
    if (distance < 0) {
      return "Scheduled time has passed";
    }

    const days = Math.floor(distance / (1000 * 60 * 60 * 24));
    const hours = Math.floor(
      (distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60)
    );
    const minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((distance % (1000 * 60)) / 1000);

    return `${days}d ${hours}h ${minutes}m ${seconds}s`;
  };
  useEffect(() => {
    setMounted(true);
  }, []);

  // useEffect(() => {
  //   fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
  //     .then((response) => {
  //       return response.json();
  //     })
  //     .then((jsonData) => {
  //       setEmojis(jsonData.emojis);
  //     });
  // }, []);
  // const queryEmojis = (query: any, callback: (emojis: any) => void) => {
  //   if (query.length === 0) return;

  //   const matches = emojis
  //     .filter((emoji: any) => {
  //       return emoji.name.indexOf(query.toLowerCase()) > -1;
  //     })
  //     .slice(0, 10);
  //   return matches.map(({ emoji }) => ({ id: emoji }));
  // };

  useEffect(() => {
    if (
      wsMsg?.broadcast_id == broadcastId &&
      wsMsg?.command == "BROADCAST_COMPLETED"
    ) {
      //   console.log("wsMsg", wsMsg);
      toast.success(wsMsg.message);
      getDetail();
    }
    if (
      wsMsg?.broadcast_id == broadcastId &&
      wsMsg?.command == "BROADCAST_PROGRESS"
    ) {
      //   console.log("wsMsg", wsMsg);
      setBroadcast({
        ...broadcast!,
        success_count: wsMsg.data.success,
        failed_count: wsMsg.data.failed,
        completed_count: wsMsg.data.completed,
      });
    }
  }, [wsMsg]);

  const update = (connections?: ConnectionModel[]) => {
    setLoading(true);

    updateBroadcast(broadcast?.id!, {
      ...broadcast!,
      connections: connections || broadcast?.connections,
      files: files,
      products: selectedProducts,
    })
      .then(() => {
        getDetail();
      })
      .catch((error) => {
        toast.error(`${error}`);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {}, [broadcast?.connections]);

  useEffect(() => {
    getConnections({ page: 1, size: 100 }).then((res: any) => {
      setConnections(res.data);
    });
  }, []);
  useEffect(() => {
    if (mounted && broadcastId) {
      setLoading(true);
      getDetail();
    }
  }, [mounted, broadcastId, page, search, size]);

  const getDetail = () => {
    getBroadcast(broadcastId!, { page, size, search })
      .then((res: any) => {
        setBroadcast(res.data);
        setPagination(res.pagination);
        setisEditable(res.data.status === "DRAFT");
        setFiles(res.data.files ?? []);
        setSelectedProducts(res.data.products ?? []);
        if (res.data.scheduled_at) {
          const intervalId = setInterval(() => {
            setCountdown(countdownToScheduledAt(res.data.scheduled_at));
          }, 1000);
          return () => clearInterval(intervalId);
        }
      })
      .catch((error) => {
        toast.error(`${error}`);
      })
      .finally(() => {
        setLoading(false);
      });
  };
  return (
    <AdminLayout>
      <div className="p-8 h-[calc(100vh-80px)] overflow-y-auto">
        <h1 className="text-2xl font-bold">Detail Broadcast</h1>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mt-4 mb-4">
          <div className="bg-white border rounded p-6 flex flex-col space-y-4">
            <div>
              <Label>Description</Label>
              {isEditable ? (
                <Textarea
                  value={broadcast?.description ?? ""}
                  onChange={(val) => {
                    setBroadcast({
                      ...broadcast!,
                      description: val.target.value,
                    });
                  }}
                />
              ) : (
                <p className="">{broadcast?.description}</p>
              )}
            </div>
            {broadcast?.template_id ? (
              <div>
                <Label>Template</Label>
                <div className="mb-4">{broadcast?.template?.title}</div>
                {(broadcast?.template?.messages ?? []).map(
                  (msg: any, index: number) => (
                    <div className="" key={index}>
                      <div className="p-4 border rounded-lg mb-4">
                        {parseMentions(msg.body, () => {})}
                      </div>
                      <div className="mb-4">
                        {(msg.files ?? []).length > 0 && (
                          <h3 className="text-lg font-semibold">Files</h3>
                        )}
                        {(msg.files ?? []).length > 0 && (
                          <div className="grid grid-cols-2 gap-2">
                            {msg.files
                              .filter((f: any) => f.mime_type.includes("image"))
                              .map((file: any, index: number) => (
                                <div key={index}>
                                  <img
                                    src={file.url}
                                    className="w-full aspect-square rounded-lg object-cover"
                                  />
                                </div>
                              ))}
                            {msg.files
                              .filter(
                                (f: any) => !f.mime_type.includes("image")
                              )
                              .map((file: any, index: number) => (
                                <div
                                  key={index}
                                  className="flex flex-col justify-center items-center text-center w-full aspect-square rounded-lg border p-4"
                                >
                                  <IoDocumentsOutline size={32} />
                                  <small className="text-center mt-4">
                                    {file?.file_name}
                                  </small>
                                </div>
                              ))}
                          </div>
                        )}
                      </div>
                      <div className="mb-4">
                        {(msg.products ?? []).length > 0 && (
                          <h3 className="text-lg font-semibold">Product</h3>
                        )}
                        {(msg.products ?? []).length > 0 && (
                          <div className="flex items-center flex-col  px-8">
                            {(msg.products[0].product_images ?? []).length >
                              0 && (
                              <img
                                src={msg.products[0].product_images![0].url}
                                alt="product"
                                className="w-32 h-32 rounded-lg"
                              />
                            )}
                            <h3 className="font-semibold mt-2 text-center">
                              {msg.products[0].name}
                            </h3>
                            <small>{money(msg.products[0].price)}</small>
                          </div>
                        )}
                      </div>
                    </div>
                  )
                )}
              </div>
            ) : (
              <>
                <div className="relative">
                  <MessageTemplateField
                    readonly={!isEditable}
                    index={0}
                    title={"Message"}
                    body={broadcast?.message ?? ""}
                    onChangeBody={(val) => {
                      setBroadcast({
                        ...broadcast!,
                        message: val,
                      });
                    }}
                    onClickEmoji={() => {}}
                    files={files}
                    onUploadFile={(file) => {
                      if (
                        (files ?? []).filter(
                          (f) => !f.mime_type.includes("image")
                        ).length === 0
                      ) {
                        // files = [file];
                        setFiles((prev) => [...prev, file]);
                      } else {
                        setFiles([
                          ...files.map((f) => {
                            if (!f.mime_type.includes("image")) {
                              return file;
                            }
                            return f;
                          }),
                        ]);
                      }
                    }}
                    onUploadImage={(file: FileModel, index?: number) => {
                      if (
                        (files ?? []).filter((f) =>
                          f.mime_type.includes("image")
                        ).length === 0
                      ) {
                        // files = [file];
                        setFiles((prev) => [...prev, file]);
                      } else {
                        setFiles([
                          ...files.map((f) => {
                            if (f.mime_type.includes("image")) {
                              return file;
                            }
                            return f;
                          }),
                        ]);
                      }
                    }}
                    onTapProduct={() => {
                      setModalProduct(true);
                    }}
                    product={selectedProducts && selectedProducts[0]}
                    onDeleteImage={(file: FileModel) => {
                      setFiles(files.filter((f) => f.id !== file.id));
                    }}
                    onDeleteFile={(file: FileModel) => {
                      setFiles(files.filter((f) => f.id !== file.id));
                    }}
                  />
                  {/* <Label>Message</Label>
                  <p className="">
                    {isEditable ? (
                      <MessageMention
                        msg={broadcast?.message ?? ""}
                        onChange={(val: any) => {
                          setBroadcast({
                            ...broadcast!,
                            message: val.target.value,
                          });
                        }}
                        onClickEmoji={() => {}}
                        onSelectEmoji={(emoji: any) => {}}
                      />
                    ) : (
                      parseMentions(broadcast?.message ?? "", (type, id) => {})
                    )}
                  </p>
                  {isEditable && (
                    <div className="absolute bottom-2 right-2 z-50">
                      <Dropdown
                        label={<BsPlusCircle />}
                        inline
                        placement="top"
                        arrowIcon={false}
                      >
                        <Dropdown.Item
                          className="flex gap-2"
                          onClick={() => {
                            fileRef.current?.click();
                          }}
                          icon={BsFileEarmark}
                        >
                          File
                        </Dropdown.Item>
                        <Dropdown.Item
                          className="flex gap-2"
                          onClick={() => {
                            getProducts({ page: 1, size: 10 }).then(
                              (res: any) => {
                                setProducts(res.data.items);
                              }
                            );
                            setModalProduct(true);
                          }}
                          icon={BsTag}
                        >
                          Product
                        </Dropdown.Item>
                      </Dropdown>
                    </div>
                  )} */}
                </div>
                {/* {((broadcast?.products ?? []).length > 0 ||
                  (broadcast?.files ?? []).length > 0) && (
                  <div className="flex flex-col gap-2 bg-gray-100 rounded-lg p-2">
                    {(broadcast?.files ?? []).length > 0 && (
                      <div className="flex flex-row gap-2 items-center cursor-pointer hover:bg-gray-200 p-2">
                        {" "}
                        <div className="rounded-full w-10 h-10 bg-gray-200 flex justify-center items-center">
                          <BsFileEarmark className="w-4 h-4 " />
                        </div>
                        <div className="flex flex-col">
                          <span className="font-semibold">
                            {(broadcast?.files ?? []).length} Files
                          </span>
                          <small>
                            {(broadcast?.files ?? [])
                              .slice(0, 3)
                              .map((e) => e.path.split("/").pop())
                              .join(", ")}
                            {(broadcast?.files ?? []).length > 3 ? "..." : ""}
                          </small>
                        </div>
                      </div>
                    )}
                    {(broadcast?.products ?? []).map(
                      (product: any, index: number) => (
                        <div
                          key={product.id}
                          className="flex flex-row gap-2 items-center cursor-pointer hover:bg-gray-200 p-2"
                          onClick={() => {
                            setSelectedProducts((prev) => [...prev, product]);
                            setModalProduct(false);
                          }}
                        >
                          {" "}
                          {(product.product_images ?? []).length !== 0 ? (
                            <img
                              src={product.product_images![0].url}
                              className="w-10 h-10 rounded-full"
                            />
                          ) : (
                            <div className="rounded-full w-10 h-10 bg-gray-200 flex justify-center items-center">
                              <BsImage className="w-4 h-4 " />
                            </div>
                          )}
                          <div className="flex flex-col">
                            <span className="font-semibold">
                              {product.name}
                            </span>
                            <small>{product.description}</small>
                          </div>
                        </div>
                      )
                    )}
                  </div>
                )}
                {(files.length > 0 || selectedProducts.length > 0) && (
                  <div className=" flex w-full bg-red-50 p-4 justify-between z-0">
                    <div className="flex flex-col">
                      {files.length > 0 && (
                        <span>{files.length} Attachments</span>
                      )}
                      {selectedProducts.length > 0 && (
                        <>
                          <span>{selectedProducts.length} Products</span>
                          <small>
                            {selectedProducts.map((e) => e.name).join(", ")}{" "}
                          </small>
                        </>
                      )}
                    </div>
                    <button
                      className="text-gray-400 hover:text-gray-600 cursor-pointer"
                      onClick={() => {
                        setFiles([]);
                        setSelectedProducts([]);
                      }}
                    >
                      <FaXmark />
                    </button>
                  </div>
                )} */}
              </>
            )}

            <div className="flex gap-2">
              {broadcast?.status === "DRAFT" && (
                <Button
                  color="success"
                  onClick={() => {
                    if ((broadcast.connections ?? []).length === 0) {
                      toast.error(
                        "Broadcast must have at least one connection"
                      );
                      return;
                    }
                    if ((broadcast.contacts ?? []).length === 0) {
                      toast.error("Broadcast must have at least one contact");
                      return;
                    }
                    setLoading(true);
                    updateBroadcast(broadcast?.id!, {
                      ...broadcast,
                      status: "READY",
                      files: files,
                      products: selectedProducts,
                    })
                      .then(() => {
                        getDetail();
                        setFiles([]);
                        setSelectedProducts([]);
                      })
                      .catch((error) => {
                        toast.error(`${error}`);
                      })
                      .finally(() => {
                        setLoading(false);
                      });
                  }}
                >
                  Ready To Send
                </Button>
              )}
              {broadcast?.status === "READY" && (
                <Button
                  color="warning"
                  onClick={() => {
                    setLoading(true);
                    updateBroadcast(broadcast?.id!, {
                      ...broadcast,
                      status: "DRAFT",
                    })
                      .then(() => {
                        getDetail();
                      })
                      .catch((error) => {
                        toast.error(`${error}`);
                      })
                      .finally(() => {
                        setLoading(false);
                      });
                  }}
                >
                  Undo
                </Button>
              )}
              {broadcast?.status === "READY" && (
                <Button
                  color="success"
                  onClick={() => {
                    if (
                      window.confirm(
                        broadcast?.scheduled_at
                          ? "Are you sure you want to send the broadcast at " +
                              moment(broadcast?.scheduled_at).format(
                                "DD MMM YYYY HH:mm"
                              )
                          : "Are you sure you want to send the broadcast?"
                      )
                    ) {
                      setLoading(true);
                      sendBroadcast(broadcast?.id!)
                        .then(() => {
                          getDetail();
                        })
                        .catch((error) => {
                          toast.error(`${error}`);
                        })
                        .finally(() => {
                          setLoading(false);
                        });
                    }
                  }}
                >
                  {broadcast?.scheduled_at ? (
                    <div>
                      Broadcast at{" "}
                      {moment(broadcast?.scheduled_at).format(
                        "DD MMM YYYY HH:mm"
                      )}
                    </div>
                  ) : (
                    "Broadcast Now"
                  )}
                </Button>
              )}
              <div className="flex flex-row gap-2">
              {broadcast?.status === "DRAFT" && (
                <Button onClick={() => update()}>Save</Button>
              )}
              {broadcast?.status === "SCHEDULED" && (
                <Button color="purple" onClick={() => {}}>
                  <GoClock className="mr-2" /> {countdown}
                </Button>
              )}
              {broadcast?.status === "SCHEDULED" &&
                scheduledTimeDistance < 0 && (
                  <Button
                    color="yellow"
                    onClick={() => {
                      setLoading(true);
                      updateBroadcast(broadcast?.id!, {
                        ...broadcast,
                        status: "DRAFT",
                      })
                        .then(() => {
                          getDetail();
                        })
                        .catch((error) => {
                          toast.error(`${error}`);
                        })
                        .finally(() => {
                          setLoading(false);
                        });
                    }}
                  >
                    <AiOutlineEdit className="mr-2" /> Edit Broadcast
                  </Button>
                )}
                </div>
            </div>
          </div>
          <div className="bg-white border rounded p-6 flex flex-col space-y-4">
            <div>
              <Label>Status</Label>
              <div className="w-fit">
                <Badge
                  color={
                    broadcast?.status === "DRAFT"
                      ? "warning"
                      : broadcast?.status === "PROCESSING"
                      ? "blue"
                      : broadcast?.status === "SCHEDULED"
                      ? "purple"
                      : broadcast?.status === "FAILED"
                      ? "danger"
                      : "success"
                  }
                >
                  {broadcast?.status}
                </Badge>
              </div>
            </div>
            {broadcast?.status !== "DRAFT" && (
              <div>
                <Label>Sequence</Label>
                <p className="">{broadcast?.group_count ?? 0}</p>
              </div>
            )}

            <div>
              <Label>Max Contact per Sequence</Label>
              {isEditable ? (
                <TextInput
                  type="number"
                  value={broadcast?.max_contacts_per_batch ?? 0}
                  onChange={(val) => {
                    setBroadcast({
                      ...broadcast!,
                      max_contacts_per_batch: Number(val.target.value),
                    });
                  }}
                />
              ) : (
                <div className="w-fit">{broadcast?.max_contacts_per_batch}</div>
              )}
            </div>
            <div>
              <Label>Sequence Delay Time (ms)</Label>
              {isEditable ? (
                <TextInput
                  type="number"
                  disabled={!isEditable}
                  value={broadcast?.sequence_delay_time ?? 0}
                  onChange={(val) => {
                    setBroadcast({
                      ...broadcast!,
                      sequence_delay_time: Number(val.target.value),
                    });
                  }}
                />
              ) : (
                <div>{broadcast?.sequence_delay_time ?? 0}s</div>
              )}
            </div>
            <div>
              <Label>Delay Time (s)</Label>
              {isEditable ? (
                <TextInput
                  type="number"
                  disabled={!isEditable}
                  value={broadcast?.delay_time ?? 0}
                  onChange={(val) => {
                    setBroadcast({
                      ...broadcast!,
                      delay_time: Number(val.target.value),
                    });
                  }}
                />
              ) : (
                <div>{broadcast?.delay_time ?? 0}s</div>
              )}
            </div>
            <div>
              <Label>Scheduled At</Label>
              {isEditable ? (
                <div className="grid grid-cols-5 gap-4">
                  <Datepicker
                    disabled={broadcast?.scheduled_at ? false : true}
                    value={
                      broadcast?.scheduled_at
                        ? moment(broadcast?.scheduled_at).toDate()
                        : null
                    }
                    onChange={(val) => {
                      setBroadcast({
                        ...broadcast!,
                        scheduled_at: val,
                      });
                    }}
                    className="col-span-2"
                  />
                  <div className="flex col-span-2">
                    <input
                      disabled={broadcast?.scheduled_at ? false : true}
                      type="time"
                      id="time"
                      style={{
                        color: broadcast?.scheduled_at ? "black" : "gray",
                      }}
                      className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                      value={moment(broadcast?.scheduled_at).format("HH:mm")}
                      onChange={(e) => {
                        setBroadcast({
                          ...broadcast!,
                          scheduled_at: moment(
                            moment(broadcast?.scheduled_at).format(
                              "YYYY-MM-DD"
                            ) +
                              " " +
                              e.target.value
                          ).toDate(),
                        });
                      }}
                    />
                    <span className="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border rounded-s-0 border-s-0 border-gray-300 rounded-e-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                      <svg
                        className="w-4 h-4 text-gray-500 dark:text-gray-400"
                        aria-hidden="true"
                        xmlns="http://www.w3.org/2000/svg"
                        fill="currentColor"
                        viewBox="0 0 24 24"
                      >
                        <path
                          fillRule="evenodd"
                          d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586V8Z"
                          clipRule="evenodd"
                        />
                      </svg>
                    </span>
                  </div>
                  <div className="flex items-center flex-col justify-center">
                    <ToggleSwitch
                      checked={broadcast?.scheduled_at ? true : false}
                      onChange={(val) => {
                        setBroadcast({
                          ...broadcast!,
                          scheduled_at: val ? new Date() : null,
                        });
                      }}
                      label="Schedule"
                    />
                  </div>
                </div>
              ) : (
                broadcast?.scheduled_at && (
                  <div>
                    <Moment className="" format="DD MMM YYYY, HH:mm">
                      {broadcast?.scheduled_at}
                    </Moment>
                  </div>
                )
              )}
            </div>
            <div>
              <Label>Connection</Label>
              <p className="">
                {isEditable ? (
                  <Select
                    options={connections.filter(
                      (item: any) =>
                        item.status === "ACTIVE" &&
                        (item.type === "whatsapp" ||
                          item.type === "whatsapp-api")
                    )}
                    value={broadcast?.connections ?? []}
                    isMulti
                    onChange={(val) => {
                      // console.log(val);

                      // setTimeout(() => {

                      // }, 500);
                      if (val.length !== 0) {
                        update(val.map((item: any) => item));
                      } else {
                        update([]);
                      }
                    }}
                    formatOptionLabel={(option) => (
                      <div className="flex items-center space-x-3">
                        <div className="flex-1 min-w-0">
                          <p className="text-sm font-medium text-gray-900 truncate">
                            {option.name}
                          </p>
                          <p className="text-xs text-gray-500 truncate">
                            {option.session_name}
                          </p>
                        </div>
                      </div>
                    )}
                  />
                ) : (
                  <div>
                    {(broadcast?.connections ?? []).map((item: any) => (
                      <div className="flex flex-col w-fit px-2 bg-amber-200 rounded-lg">
                        <div className="font-semibold">{item.name}</div>
                        <small className="">{item.session_name}</small>
                      </div>
                    ))}
                  </div>
                )}
              </p>
            </div>
            <div>
              <Label>Total Contact</Label>
              <p className="">{broadcast?.contact_count ?? 0}</p>
            </div>
            <div>
              <Label>Progress</Label>
              <p className="">
                {money(
                  ((broadcast?.completed_count ?? 0) /
                    (broadcast?.contact_count ?? 0)) *
                    100
                )}
                %
              </p>
            </div>

            {broadcast?.status === "COMPLETED" && (
              <Chart
                chartType="PieChart"
                width="100%"
                height="300px"
                data={[
                  ["Status", "Count"],
                  ["Succeed", broadcast?.success_count ?? 0],
                  ["Failed", broadcast?.failed_count ?? 0],
                ]}
                options={{
                  colors: ["#10b981", "#ef4444", "#f59e0b"],
                  pieHole: 0.4,
                  legend: { position: "bottom" },
                  is3D: true,
                }}
              />
            )}
            {broadcast?.status === "PROCESSING" && (
              <Chart
                chartType="PieChart"
                width="100%"
                height="300px"
                data={[
                  ["Status", "Count"],
                  ["Completed", broadcast?.completed_count ?? 0],
                  [
                    "Uncompleted",
                    (broadcast?.contact_count ?? 0) -
                      (broadcast?.completed_count ?? 0),
                  ],
                ]}
                options={{
                  colors: ["#10b981", "#ef4444", "#f59e0b"],
                  pieHole: 0.4,
                  legend: { position: "bottom" },
                  is3D: true,
                }}
              />
            )}
          </div>
        </div>
        <div className="bg-white border rounded p-6 flex flex-col space-y-4">
          <div className="flex items-center justify-between">
            <h3 className="text-lg font-bold">Contact</h3>
            <div className="flex gap-2">
              {selectedBroadcastContacts.length > 0 &&
                broadcast?.status === "DRAFT" && (
                  <Button
                    color="red"
                    size="sm"
                    onClick={() => {
                      if (
                        window.confirm(
                          `Are you sure you want to delete ${selectedBroadcastContacts.length} contacts ?`
                        )
                      ) {
                        deleteContactBroadcast(broadcastId!, {
                          contact_ids: selectedBroadcastContacts.map(
                            (item) => item.id
                          ),
                        }).then(() => {
                          getDetail();
                          setSelectedBroadcastContacts([]);
                        });
                      }
                    }}
                  >
                    Delete
                  </Button>
                )}
              {broadcast?.status === "DRAFT" && (
                <Button
                  color="gray"
                  size="sm"
                  onClick={() => {
                    countContactByTag().then((v: any) => {
                      setTags(v.data);
                    });
                    getContacts({ page: 1, size: 10 }).then((v: any) => {
                      setContacts(v.data.items);
                    });
                    setShowModal(true);
                  }}
                >
                  + Contact
                </Button>
              )}
            </div>
          </div>
          <div className="relative w-full max-w-[300px] mr-6 focus-within:text-purple-500">
            <div className="absolute inset-y-0 left-0 flex items-center pl-3">
              <HiMagnifyingGlass />
            </div>
            <input
              type="text"
              className="w-full py-2 pl-10 text-sm text-gray-700 bg-white border border-gray-300 rounded-2xl shadow-sm focus:outline-none focus:ring focus:ring-indigo-200 focus:border-indigo-500"
              placeholder="Search"
              value={search}
              onChange={(e) => {
                setSearch(e.target.value);
              }}
            />
          </div>
          <Table>
            <Table.Head>
              <Table.HeadCell>
                <Checkbox
                  disabled={!isEditable}
                  checked={
                    selectedBroadcastContacts.length ===
                    (broadcast?.contacts ?? []).length
                  }
                  onChange={(e) => {
                    if (e.target.checked) {
                      setSelectedBroadcastContacts(broadcast?.contacts ?? []);
                    } else {
                      setSelectedBroadcastContacts([]);
                    }
                  }}
                  className="mr-2"
                />
                Name
              </Table.HeadCell>
              {/* <Table.HeadCell>Email</Table.HeadCell> */}
              <Table.HeadCell>Phone</Table.HeadCell>
              {/* <Table.HeadCell>Address</Table.HeadCell> */}
              <Table.HeadCell>Tag</Table.HeadCell>
              <Table.HeadCell>Data</Table.HeadCell>
              <Table.HeadCell>Info</Table.HeadCell>
              <Table.HeadCell></Table.HeadCell>
            </Table.Head>

            <Table.Body className="divide-y">
              {(broadcast?.contacts ?? []).length === 0 && (
                <Table.Row>
                  <Table.Cell colSpan={5} className="text-center">
                    No contacts found.
                  </Table.Cell>
                </Table.Row>
              )}
              {(broadcast?.contacts ?? []).map((contact) => (
                <Table.Row
                  key={contact.id}
                  className="bg-white dark:border-gray-700 dark:bg-gray-800"
                >
                  <Table.Cell
                    className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                    onClick={() => {}}
                  >
                    <div className="flex gap-2">
                      <Checkbox
                        disabled={!isEditable}
                        checked={selectedBroadcastContacts
                          .map((contact) => contact.id)
                          .includes(contact.id)}
                        onChange={(e) => {
                          if (e.target.checked) {
                            setSelectedBroadcastContacts([
                              ...selectedBroadcastContacts,
                              contact,
                            ]);
                          } else {
                            setSelectedBroadcastContacts(
                              selectedBroadcastContacts.filter(
                                (c) => c.id !== contact.id
                              )
                            );
                          }
                        }}
                      />
                      <div className="flex flex-col">
                        {contact.name}
                        {(contact.tags ?? []).length > 0 && (
                          <div className="flex flex-wrap gap-2">
                            {contact.tags?.map((tag) => (
                              <span
                                className="px-2  text-[8pt] font-semibold text-gray-900 bg-gray-100 rounded dark:bg-gray-700 dark:text-gray-100"
                                key={tag.id}
                                style={{
                                  color: getContrastColor(tag.color),
                                  backgroundColor: tag.color,
                                }}
                              >
                                {tag.name}
                              </span>
                            ))}
                          </div>
                        )}
                      </div>
                    </div>
                  </Table.Cell>
                  {/* <Table.Cell>{contact.email}</Table.Cell> */}
                  <Table.Cell>{contact.phone}</Table.Cell>
                  {/* <Table.Cell>{contact.address}</Table.Cell> */}
                  <Table.Cell>
                    {(contact.tags ?? []).map((tag) => tag.name).join(", ")}
                  </Table.Cell>
                  <Table.Cell>
                    {Object.keys(contact.custom_data).map(
                      (key) => `${key}: ${contact.custom_data[key]}`
                    )}
                  </Table.Cell>
                  <Table.Cell className="w-32">
                    {(broadcast?.status === "COMPLETED" ||
                      broadcast?.status === "PROCESSING") && (
                      <div className="flex gap-2">
                        <ul>
                          <li className="flex gap-2 w-full justify-between">
                            <span>Completed</span>
                            {contact.is_completed && (
                              <BsCheck2Circle className="text-green-500" />
                            )}
                          </li>
                          <li className="flex gap-2 w-full justify-between">
                            <span>Success</span>
                            {contact.is_success ? (
                              <BsCheck2Circle className="text-green-500" />
                            ) : (
                              <Tooltip
                                content={
                                  (contact.data?.logs ?? []).length > 0 &&
                                  contact.data?.logs?.[0]?.error_message
                                }
                              >
                                <FaXmark className="text-red-500" />
                              </Tooltip>
                            )}
                          </li>

                          <li className="flex gap-2 w-full justify-between">
                            <span>Retries</span>
                            {contact.data?.retry?.attempt ?? 0}
                          </li>
                        </ul>
                      </div>
                    )}
                  </Table.Cell>
                  <Table.Cell>
                    {isEditable && (
                      <a
                        href="#"
                        className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                        onClick={(e) => {
                          e.preventDefault();
                          if (
                            window.confirm(
                              `Are you sure you want to delete contact ${contact.name}?`
                            )
                          ) {
                            // deleteContact(contact?.id!).then(() => {
                            //   getAllContacts();
                            // });

                            deleteContactBroadcast(broadcastId!, {
                              contact_ids: [contact.id],
                            }).then(() => {
                              getDetail();
                            });
                          }
                        }}
                      >
                        Delete
                      </a>
                    )}
                  </Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
          <Pagination
            className="mt-4"
            currentPage={page}
            totalPages={pagination?.total_pages ?? 0}
            onPageChange={(val) => {
              setPage(val);
            }}
            showIcons
          />
        </div>
      </div>
      <Modal
        show={showModal}
        onClose={() => setShowModal(false)}
        title="Add Contact"
        size="4xl"
      >
        <Modal.Header>Add Contact</Modal.Header>
        <Modal.Body>
          <Tabs>
            <TabItem title="Contact">
              <div className="mb-4 flex justify-end">
                <div className="relative w-full max-w-[300px] mr-6 focus-within:text-purple-500">
                  <div className="absolute inset-y-0 left-0 flex items-center pl-3">
                    <HiMagnifyingGlass />
                  </div>
                  <input
                    type="text"
                    className="w-full py-2 pl-10 text-sm text-gray-700 bg-white border border-gray-300 rounded-2xl shadow-sm focus:outline-none focus:ring focus:ring-indigo-200 focus:border-indigo-500"
                    placeholder="Search"
                    onChange={(e) => {
                      getContacts({
                        page: 1,
                        size: 10,
                        search: e.target.value,
                      }).then((v: any) => {
                        setContacts(v.data.items);
                      });
                    }}
                  />
                </div>
              </div>
              <table className="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                  <tr>
                    <th
                      scope="col"
                      className="px-3 py-3"
                      style={{ width: "5%" }}
                    ></th>
                    <th scope="col" className="px-6 py-3">
                      Name
                    </th>
                    <th scope="col" className="px-6 py-3">
                      Email
                    </th>
                    <th scope="col" className="px-6 py-3">
                      Phone
                    </th>
                    <th scope="col" className="px-6 py-3">
                      Address
                    </th>
                    <th scope="col" className="px-6 py-3">
                      Position
                    </th>
                  </tr>
                </thead>
                <tbody>
                  {contacts.map((contact) => (
                    <tr
                      key={contact.id}
                      className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
                    >
                      <td className="px-3 py-4">
                        <Checkbox
                          checked={selectedContacts
                            .map((contact) => contact.id)
                            .includes(contact.id)}
                          onChange={(e) => {
                            if (e.target.checked) {
                              setSelectedContacts([
                                ...selectedContacts,
                                contact,
                              ]);
                            } else {
                              setSelectedContacts(
                                selectedContacts.filter(
                                  (c) => c.id !== contact.id
                                )
                              );
                            }
                          }}
                        />
                      </td>
                      <td className="px-6 py-4">{contact.name}</td>
                      <td className="px-6 py-4">{contact.email}</td>
                      <td className="px-6 py-4">{contact.phone}</td>
                      <td className="px-6 py-4">{contact.address}</td>
                      <td className="px-6 py-4">
                        {contact.contact_person_position}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </TabItem>
            <TabItem title="Tag">
              <ul>
                {(tags ?? []).map((item: any) => (
                  <li className="flex items-center gap-2 py-2">
                    <Checkbox
                      checked={selectedTags
                        .map((tag) => tag.id)
                        .includes(item.id)}
                      onChange={(e) => {
                        if (e.target.checked) {
                          setSelectedTags([...selectedTags, item]);
                        } else {
                          setSelectedTags(
                            selectedTags.filter((t) => t.id !== item.id)
                          );
                        }
                      }}
                    />
                    <p>{item.name}</p>
                    <p>( {item.count} )</p>
                  </li>
                ))}
              </ul>
            </TabItem>
            <TabItem title="From File">
              <div>
                <FileInput
                  id="file"
                  name="file"
                  onChange={async (e) => {
                    const file = e.target.files![0];
                    if (file) {
                      try {
                        setLoading(true);
                        const resp: any = await uploadFile(
                          file,
                          {},
                          (v: any) => {
                            console.log(v);
                          }
                        );
                        let res2: any = await addContactFromFileBroadcast(
                          broadcast!.id!,
                          {
                            file_url: resp.data.url,
                          }
                        );

                        console.log(res2);
                        setShowModal(false);
                        getDetail();
                      } catch (error) {
                        console.error(error);
                      } finally {
                        setLoading(false);
                      }
                    }
                  }}
                />
              </div>
            </TabItem>
          </Tabs>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex justify-end w-full gap-2">
            <Button
              onClick={() => {
                setLoading(true);
                addContactBroadcast(broadcast!.id!, {
                  tag_ids: selectedTags.map((tag) => tag.id),
                  contact_ids: selectedContacts.map((contact) => contact.id),
                })
                  .then((v: any) => {
                    getDetail();
                    toast.success("Successfully added contact");
                    setShowModal(false);
                  })
                  .catch((e: any) => {
                    toast.error(`${e.message}`);
                  })
                  .finally(() => {
                    setLoading(false);
                  });
              }}
            >
              Save
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
      <input
        multiple
        accept=".png, .jpg, .jpeg, .doc, .docx, .xls, .xlsx, .pdf"
        type="file"
        name="file"
        id=""
        ref={fileRef}
        className="hidden"
        onChange={async (e) => {
          if ((e.target.files ?? []).length > 0) {
            for (
              let index = 0;
              index < (e.target.files ?? []).length;
              index++
            ) {
              const element = (e.target.files ?? [])[index];
              let resp: any = await uploadFile(element, {}, console.log);
              setFiles((prev) => [...prev, resp.data]);
            }
          }
        }}
      />
      <Modal show={modalProduct} onClose={() => setModalProduct(false)}>
        <Modal.Header>Product</Modal.Header>
        <Modal.Body>
          <div className="relative w-full mb-8 mr-6 focus-within:text-purple-500">
            <div className="absolute inset-y-0 left-0 flex items-center pl-3">
              <HiMagnifyingGlass />
            </div>
            <input
              type="text"
              className="w-full py-2 pl-10 text-sm text-gray-700 bg-white border border-gray-300 rounded-2xl shadow-sm focus:outline-none focus:ring focus:ring-indigo-200 focus:border-indigo-500"
              placeholder="Search"
              onChange={(e) => {
                getProducts({
                  page: 1,
                  search: e.target.value,
                  size: 10,
                }).then((res: any) => {
                  setProducts(res.data.items);
                });
              }}
            />
          </div>
          {products.length === 0 && (
            <div className="text-center">No product found.</div>
          )}
          <div className="flex flex-col gap-2">
            {products.map((product) => (
              <div
                key={product.id}
                className="flex flex-row gap-2 items-center cursor-pointer hover:bg-gray-100 p-2"
                onClick={() => {
                  setSelectedProducts((prev) => [...prev, product]);
                  setModalProduct(false);
                }}
              >
                {" "}
                {product.product_images?.length !== 0 ? (
                  <img
                    src={product.product_images![0].url}
                    className="w-10 h-10 rounded-full"
                  />
                ) : (
                  <div className="rounded-full w-10 h-10 bg-gray-200 flex justify-center items-center">
                    <BsImage className="w-4 h-4 " />
                  </div>
                )}
                <div className="flex flex-col">
                  <span className="font-semibold">{product.name}</span>
                  <small>{product.description}</small>
                </div>
              </div>
            ))}
          </div>
        </Modal.Body>
      </Modal>
      <ModalProductList
        show={modalProduct}
        setShow={setModalProduct}
        selectProduct={(val) => {
          setSelectedProducts((prev) => [val]);
        }}
      />
    </AdminLayout>
  );
};
export default BroadcastDetail;

const emojiStyle = {
  control: {
    fontSize: 16,
    lineHeight: 1.2,
    minHeight: 160,
  },

  highlighter: {
    padding: 9,
    border: "1px solid transparent",
  },

  input: {
    fontSize: 16,
    lineHeight: 1.2,
    padding: 9,
    border: "1px solid silver",
    borderRadius: 10,
  },

  suggestions: {
    list: {
      backgroundColor: "white",
      border: "1px solid rgba(0,0,0,0.15)",
      fontSize: 16,
    },

    item: {
      padding: "5px 15px",
      borderBottom: "1px solid rgba(0,0,0,0.15)",

      "&focused": {
        backgroundColor: "#cee4e5",
      },
    },
  },
};
