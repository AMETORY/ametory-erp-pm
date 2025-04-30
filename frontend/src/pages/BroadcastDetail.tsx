import { useContext, useEffect, useState, type FC } from "react";
import { useParams } from "react-router-dom";
import AdminLayout from "../components/layouts/admin";
import {
  addContactBroadcast,
  deleteContactBroadcast,
  getBroadcast,
  sendBroadcast,
  updateBroadcast,
} from "../services/api/broadcastApi";
import { BroadcastModel } from "../models/broadcast";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import {
  Badge,
  Button,
  Checkbox,
  Datepicker,
  Label,
  Modal,
  Pagination,
  TabItem,
  Table,
  Tabs,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import { parseMentions } from "../utils/helper-ui";
import { Mention, MentionsInput } from "react-mentions";
import Moment from "react-moment";
import moment from "moment";
import Select from "react-select";
import { ConnectionModel } from "../models/connection";
import { getConnections } from "../services/api/connectionApi";
import { countContactByTag, getContacts } from "../services/api/contactApi";
import { TagModel } from "../models/tag";
import { getContrastColor, getPagination, money } from "../utils/helper";
import { ContactModel } from "../models/contact";
import { HiMagnifyingGlass } from "react-icons/hi2";
import { SearchContext } from "../contexts/SearchContext";
import { PaginationResponse } from "../objects/pagination";
import { BsCheck2Circle, BsInfoCircle } from "react-icons/bs";
import { FaXmark } from "react-icons/fa6";
import { WebsocketContext } from "../contexts/WebsocketContext";
import Chart from "react-google-charts";

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
  const [selectedBroadcastContacts, setSelectedBroadcastContacts] = useState<
    ContactModel[]
  >([]);

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    fetch(
      "https://gist.githubusercontent.com/oliveratgithub/0bf11a9aff0d6da7b46f1490f86a71eb/raw/d8e4b78cfe66862cf3809443c1dba017f37b61db/emojis.json"
    )
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);
  const queryEmojis = (query: any, callback: (emojis: any) => void) => {
    if (query.length === 0) return;

    const matches = emojis
      .filter((emoji: any) => {
        return emoji.name.indexOf(query.toLowerCase()) > -1;
      })
      .slice(0, 10);
    return matches.map(({ emoji }) => ({ id: emoji }));
  };

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
            <div>
              <Label>Message</Label>
              <p className="">
                {isEditable ? (
                  <MentionsInput
                    value={broadcast?.message ?? ""}
                    onChange={(val) => {
                      setBroadcast({
                        ...broadcast!,
                        message: val.target.value,
                      });
                    }}
                    style={emojiStyle}
                    placeholder={
                      "Press ':' for emojis, and template using '@' and shift+enter to send"
                    }
                    autoFocus
                  >
                    <Mention
                      trigger="@"
                      data={[
                        { id: "{{user}}", display: "Full Name" },
                        { id: "{{phone}}", display: "Phone Number" },
                        { id: "{{agent}}", display: "Agent Name" },
                      ]}
                      style={{
                        backgroundColor: "#cee4e5",
                      }}
                      appendSpaceOnAdd
                    />
                    <Mention
                      trigger=":"
                      markup="__id__"
                      regex={neverMatchingRegex}
                      data={queryEmojis}
                    />
                  </MentionsInput>
                ) : (
                  parseMentions(broadcast?.message ?? "", (type, id) => {})
                )}
              </p>
            </div>
            <div>
              <Label>Status</Label>
              <div className="w-fit">
                <Badge
                  color={
                    broadcast?.status === "DRAFT"
                      ? "warning"
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
            <div className="flex gap-2">
              {broadcast?.status === "DRAFT" && (
                <Button
                  color="success"
                  onClick={() => {
                    setLoading(true);
                    updateBroadcast(broadcast?.id!, {
                      ...broadcast,
                      status: "READY",
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
              {broadcast?.status === "DRAFT" && (
                <Button
                  onClick={() => {
                    setLoading(true);
                    updateBroadcast(broadcast?.id!, broadcast)
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
                  Save
                </Button>
              )}
            </div>
          </div>
          <div className="bg-white border rounded p-6 flex flex-col space-y-4">
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
                      (item: any) => item.status === "ACTIVE"
                    )}
                    value={broadcast?.connections ?? []}
                    isMulti
                    onChange={(val) => {
                      console.log(val);
                      setBroadcast({
                        ...broadcast!,
                        connections: val.map((item: any) => item),
                      });
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
              <p className="">{money((broadcast?.completed_count ?? 0)/(broadcast?.contact_count ?? 0) * 100)}%</p>
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
              <Table.HeadCell>Name</Table.HeadCell>
              <Table.HeadCell>Email</Table.HeadCell>
              <Table.HeadCell>Phone</Table.HeadCell>
              <Table.HeadCell>Address</Table.HeadCell>
              <Table.HeadCell>Position</Table.HeadCell>
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
                  <Table.Cell>{contact.email}</Table.Cell>
                  <Table.Cell>{contact.phone}</Table.Cell>
                  <Table.Cell>{contact.address}</Table.Cell>
                  <Table.Cell>{contact.contact_person_position}</Table.Cell>
                  <Table.Cell className="w-32">
                    {broadcast?.status === "COMPLETED" && (
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
                              <FaXmark className="text-red-500" />
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
                {tags.map((item: any) => (
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
