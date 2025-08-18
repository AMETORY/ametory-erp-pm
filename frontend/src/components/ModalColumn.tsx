import { Button, Label, Modal, TextInput, ToggleSwitch } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import toast from "react-hot-toast";
import { BsCamera, BsCheck2Circle, BsTrash } from "react-icons/bs";
import { GoPaperclip } from "react-icons/go";
import { HiOutlineDocumentAdd } from "react-icons/hi";
import { IoDocumentsOutline } from "react-icons/io5";
import { Mention, MentionsInput } from "react-mentions";
import Select from "react-select";
import { ColumnActionModel, ColumnModel } from "../models/column";
import { FileModel } from "../models/file";
import { deleteFile, uploadFile } from "../services/api/commonApi";
import {
  addNewColumnAction,
  deleteColumnAction,
  editColumnAction,
  getColumn,
  updateColumn,
} from "../services/api/projectApi";
import { parseMentions } from "../utils/helper-ui";
import MessageTemplateField from "./MessageTemplateField";
import moment from "moment";
const neverMatchingRegex = /($a)/;

interface ModalColumnProps {
  projectId: string;
  column: ColumnModel;
  showModal: boolean;
  setShowModal: (val: boolean) => void;
  onChangeColumn: (column: ColumnModel) => void;
  onAddAction: (column: ColumnModel) => void;
}

const ModalColumn: FC<ModalColumnProps> = ({
  projectId,
  column,
  showModal,
  setShowModal,
  onChangeColumn,
  onAddAction,
}) => {
  const [selectedAction, setSelectedAction] = useState<ColumnActionModel>();
  const [actions, setActions] = useState<ColumnActionModel[]>([]);
  const [emojis, setEmojis] = useState([]);
  let triggers = [
    {
      label: "IDLE",
      value: "IDLE",
    },
    {
      label: "Move In",
      value: "MOVE_IN",
    },
    {
      label: "Move Out",
      value: "MOVE_OUT",
    },
  ];
  let idleTimeType = [
    {
      label: "Days",
      value: "days",
    },
    {
      label: "Hours",
      value: "hours",
    },
    {
      label: "Minutes",
      value: "minutes",
    },
  ];

  useEffect(() => {
    if (column) {
      setActions(column.actions ?? []);
    }
  }, [column]);

  useEffect(() => {
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
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
  return (
    <>
      <Modal size="7xl" show={showModal} onClose={() => setShowModal(false)}>
        <Modal.Header>Edit Column</Modal.Header>
        <Modal.Body className="space-y-6">
          <div>
            <div className="mb-2 block">
              <Label htmlFor="icon" value=" Icon" />
            </div>
            <TextInput
              id="icon"
              value={column.icon}
              onChange={(e) =>
                onChangeColumn({ ...column, icon: e.target.value })
              }
              placeholder=" Icon"
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Name" />
            </div>
            <TextInput
              id="name"
              value={column.name}
              onChange={(e) =>
                onChangeColumn({ ...column, name: e.target.value })
              }
              placeholder="Name"
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="color" value="Color" />
            </div>
            <input
              id="color"
              type="color"
              value={column.color}
              onChange={(e) =>
                onChangeColumn({ ...column, color: e.target.value })
              }
            />
          </div>
          <div>
            <div className="mb-2 block">
              <div className="flex justify-between items-center">
                <Label htmlFor="action" value="Actions" />
                <Button
                  size="xs"
                  color="gray"
                  onClick={() => {
                    addNewColumnAction(projectId!, column!.id!, {
                      name: "New Action",
                      action: "send_whatsapp_message",
                      action_trigger: "IDLE",
                      action_data: {
                        message: "Hello",
                        idle_time: 1,
                        idle_time_type: "days",
                      },
                    }).then(() => {
                      getColumn(projectId!, column!.id!).then((resp: any) => {
                        onChangeColumn(resp.data);
                        onAddAction(resp.data);
                      });
                    });
                  }}
                >
                  + Action
                </Button>
              </div>
              <table className="w-full text-sm mt-4">
                <thead>
                  <tr className="bg-gray-50">
                    <th className="text-left px-2 py-2 border">Name</th>
                    <th className="text-left px-2 py-2 border">
                      Action Trigger
                    </th>
                    <th className="text-left px-2 py-2 border">Action</th>
                    <th className="text-left px-2 py-2 border">Action Data</th>
                    <th className="text-left px-2 py-2 border">Files</th>
                    <th className="text-left px-2 py-2 border"></th>
                  </tr>
                </thead>
                <tbody>
                  {actions.map((action) => (
                    <tr key={action.id}>
                      <td className="px-2 py-2 border">{action.name}</td>
                      <td className="px-2 py-2 border">
                        {action.action_trigger.replaceAll("_", " ")}
                      </td>
                      <td className="px-2 py-2 border">
                        {action.action.toUpperCase().replaceAll("_", " ")}
                      </td>

                      <td className="px-2 py-2 border">
                        {Object.keys(action.action_data).map((key) => (
                          <div key={key}>
                            {(action.action_trigger == "IDLE" ||
                              key == "message") && (
                              <Label>
                                {key
                                  .replaceAll("_", " ")
                                  .replace(/^\w/, (c) => c.toUpperCase())}
                              </Label>
                            )}
                            {(action.action_trigger == "IDLE" ||
                              key == "message") && (
                              <div>
                                {key == "message"
                                  ? parseMentions(
                                      action.action_data[key],
                                      (type, id) => {}
                                    )
                                  : action.action_data[key]}
                              </div>
                            )}
                          </div>
                        ))}
                        {action.action_trigger == "IDLE" && (
                          <>
                            <div>
                              <Label>Hour</Label>
                              <div>{action?.action_hour}</div>
                            </div>
                            <div>
                              <Label>Action Status</Label>
                              <div>{action?.action_status}</div>
                            </div>
                          </>
                        )}
                      </td>
                      <td className="px-2 py-2 border">
                        <div className="grid grid-cols-2 ">
                          {action.files?.map((file) => (
                            <div key={file.id}>
                              {file.mime_type.includes("image") ? (
                                <img
                                  src={file.url}
                                  className="w-16 h-16 rounded-lg border object-cover"
                                />
                              ) : (
                                <div className="w-16 h-16 rounded-lg border flex justify-center items-center ">
                                  <GoPaperclip />
                                </div>
                              )}
                            </div>
                          ))}
                        </div>
                      </td>
                      <td className="px-2 py-2 border">
                        {action.status == "ACTIVE" ? (
                          <BsCheck2Circle
                            className="text-green-500"
                            size={20}
                          />
                        ) : (
                          ""
                        )}
                        <a
                          href="#"
                          className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                          onClick={() => setSelectedAction(action)}
                        >
                          Edit
                        </a>
                        <a
                          href="#"
                          className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                          onClick={(el) => {
                            el.preventDefault();
                            if (
                              window.confirm(
                                `Are you sure you want to delete  ${action.name}?`
                              )
                            ) {
                              deleteColumnAction(
                                projectId!,
                                column!.id!,
                                action!.id!
                              ).then(() => {
                                getColumn(projectId!, column!.id!).then(
                                  (resp: any) => {
                                    onChangeColumn(resp.data);
                                    onAddAction(resp.data);
                                  }
                                );
                              });
                            }
                          }}
                        >
                          Delete
                        </a>

                        {action.run_once? (
                          <div className="text-gray-600 text-xs">Run Once</div>
                        ) : (
                          ""
                        )}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button
            onClick={() => {
              updateColumn(projectId!, {
                ...column,
              })
                .then(() => setShowModal(false))
                .catch(toast.error);
            }}
          >
            Save
          </Button>
          <Button color="gray" onClick={() => setShowModal(false)}>
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
      <Modal
        size="5xl"
        dismissible
        show={!!selectedAction}
        onClose={() => setSelectedAction(undefined)}
      >
        <Modal.Header>Edit Action</Modal.Header>
        <Modal.Body className="space-y-6">
          <div className="flex flex-col space-y-4 pb-32">
            <div>
              <Label htmlFor="name" value="Name" />
              <TextInput
                id="name"
                value={selectedAction?.name}
                onChange={(e) =>
                  setSelectedAction({
                    ...selectedAction!,
                    name: e.target.value,
                  })
                }
                placeholder="Name"
              />
            </div>
            <div>
              <Label htmlFor="trigger" value="Trigger" />
              <Select
                options={triggers}
                value={triggers.find(
                  (e) => e.value === selectedAction?.action_trigger
                )}
                onChange={(e) =>
                  setSelectedAction({
                    ...selectedAction!,
                    action_trigger: e!.value!,
                  })
                }
              />
            </div>
            <div>
              <Label htmlFor="action" value="Action" />
              {/* <TextInput
                id="action"
                value={selectedAction?.action}
                onChange={(e) =>
                  setSelectedAction({
                    ...selectedAction!,
                    name: e.target.value,
                  })
                }
                placeholder="Action"
                readOnly
              /> */}
              <p>{selectedAction?.action.toUpperCase().replaceAll("_", " ")}</p>
            </div>
            <div>
              <ToggleSwitch
                label={
                  selectedAction?.status === "ACTIVE" ? "Active" : "Inactive"
                }
                onChange={(e) => {
                  setSelectedAction({
                    ...selectedAction!,
                    status: e ? "ACTIVE" : "INACTIVE",
                  });
                }}
                checked={selectedAction?.status === "ACTIVE"}
              />
            </div>
            {selectedAction?.action === "send_whatsapp_message" && (
              <>
                <div>
                  <MessageTemplateField
                    disableProduct
                    disableInteractive
                    index={0}
                    title={"Message"}
                    body={selectedAction?.action_data?.message ?? ""}
                    onChangeBody={(val) => {
                      setSelectedAction({
                        ...selectedAction!,
                        action_data: {
                          ...selectedAction!.action_data!,
                          message: val,
                        },
                      });
                    }}
                    onClickEmoji={function (): void {}}
                    files={selectedAction?.files ?? []}
                    onUploadFile={function (
                      file: FileModel,
                      index?: number
                    ): void {
                      let files = selectedAction?.files ?? [];
                      if (
                        files.filter((f) => !f.mime_type.includes("image"))
                          .length === 0
                      ) {
                        files = [...files, file];
                      } else {
                        files = files.map((f) => {
                          if (!f.mime_type.includes("image")) {
                            return file;
                          }
                          return f;
                        });
                      }
                      setSelectedAction({
                        ...selectedAction,
                        files: files,
                      });
                    }}
                    onUploadImage={function (
                      file: FileModel,
                      index?: number
                    ): void {
                      let files = selectedAction?.files ?? [];
                      if (
                        files.filter((f) => f.mime_type.includes("image"))
                          .length === 0
                      ) {
                        files = [...files, file];
                      } else {
                        files = files.map((f) => {
                          if (f.mime_type.includes("image")) {
                            return file;
                          }
                          return f;
                        });
                      }
                      setSelectedAction({
                        ...selectedAction,
                        files: files,
                      });
                    }}
                    onDeleteFile={(file: FileModel) => {
                      const confirmFirst = window.confirm(
                        "Are you sure you want to delete this file?"
                      );
                      if (!confirmFirst) return;

                      deleteFile(file?.id!).then(() => {
                        setSelectedAction({
                          ...selectedAction!,
                          files: (selectedAction?.files ?? []).filter(
                            (f) => f.id !== file.id
                          ),
                        });
                      });
                    }}
                    onDeleteImage={(file: FileModel) => {
                      const confirmFirst = window.confirm(
                        "Are you sure you want to delete this image?"
                      );
                      if (!confirmFirst) return;

                      deleteFile(file?.id!).then(() => {
                        setSelectedAction({
                          ...selectedAction!,
                          files: (selectedAction?.files ?? []).filter(
                            (f) => f.id !== file.id
                          ),
                        });
                      });
                    }}
                  />

                  {/* <Label htmlFor="action_data" value="Message" />
                  <MentionsInput
                    value={selectedAction?.action_data?.message}
                    onChange={(val) => {
                      setSelectedAction({
                        ...selectedAction!,
                        action_data: {
                          ...selectedAction!.action_data!,
                          message: val.target.value,
                        },
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
                  </MentionsInput> */}
                </div>
                {/* <div className="mt-8">
                  <h4 className="font-semibold">Files</h4>
                  <div className="grid grid-cols-2 gap-4">
                    <div className="relative h-fit">
                      <div
                        className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100 border h-[300px]"
                        onClick={() => {
                          document.getElementById(`image-action`)?.click();
                        }}
                      >
                        {(selectedAction?.files ?? []).filter((f) =>
                          f.mime_type.includes("image")
                        ).length === 0 ? (
                          <div className="flex flex-col items-center">
                            <span>Add Photo to message</span>
                            <BsCamera />
                          </div>
                        ) : (
                          <img
                            className="w-32 h-32 object-cover"
                            src={
                              (selectedAction?.files ?? []).find((f) =>
                                f.mime_type.includes("image")
                              )?.url
                            }
                          />
                        )}
                        <input
                          type="file"
                          className="hidden"
                          accept="image/*"
                          id={`image-action`}
                          onChange={(e) => {
                            const file = e.target.files?.[0];
                            if (file) {
                              uploadFile(file, {}, () => {}).then(
                                (resp: any) => {
                                  let files = selectedAction?.files ?? [];
                                  if (
                                    files.filter((f) =>
                                      f.mime_type.includes("image")
                                    ).length === 0
                                  ) {
                                    files = [...files, resp.data];
                                  } else {
                                    files = files.map((f) => {
                                      if (f.mime_type.includes("image")) {
                                        return resp.data;
                                      }
                                      return f;
                                    });
                                  }
                                  setSelectedAction({
                                    ...selectedAction,
                                    files: files,
                                  });
                                }
                              );
                            }
                          }}
                        />
                      </div>
                      <BsTrash
                        size={20}
                        className="absolute bottom-2 right-2 cursor-pointer text-red-400 hover:text-red-600"
                        onClick={() => {
                          if (
                            (selectedAction?.files ?? []).find((f) =>
                              f.mime_type.includes("image")
                            )
                          ) {
                            const confirmFirst = window.confirm(
                              "Are you sure you want to delete this image?"
                            );
                            if (!confirmFirst) return;

                            deleteFile(
                              (selectedAction?.files ?? []).find((f) =>
                                f.mime_type.includes("image")
                              )?.id!
                            ).then(() => {
                              setSelectedAction({
                                ...selectedAction!,
                                files: (selectedAction?.files ?? []).filter(
                                  (f) => !f.mime_type.includes("image")
                                ),
                              });
                            });
                          }
                        }}
                      />
                    </div>
                    <div className="relative h-fit">
                      <div
                        className="flex flex-col justify-center items-center p-16 rounded-lg bg-white cursor-pointer transition duration-300 ease-in-out hover:bg-gray-100 border h-[300px]"
                        onClick={() => {
                          document.getElementById(`image-action-file`)?.click();
                        }}
                      >
                        {(selectedAction?.files ?? []).filter(
                          (f) => !f.mime_type.includes("image")
                        ).length === 0 ? (
                          <div className="flex flex-col items-center">
                            <span>Add File to message</span>
                            <HiOutlineDocumentAdd size={32} />
                          </div>
                        ) : (
                          // <IoAttach className="rotate-[30deg]" size={32}/>
                          <div className="flex items-center flex-col px-8">
                            <IoDocumentsOutline size={32} />
                            <small className="text-center mt-4">
                              {
                                (selectedAction?.files ?? []).find(
                                  (f) => !f.mime_type.includes("image")
                                )?.file_name
                              }
                            </small>
                          </div>
                        )}
                        <input
                          type="file"
                          className="hidden"
                          accept=".doc,.docx,.pdf,.xls,.xlsx,.txt"
                          id={`image-action-file`}
                          onChange={(e) => {
                            const file = e.target.files?.[0];
                            if (file) {
                              uploadFile(file, {}, () => {}).then(
                                (resp: any) => {
                                  if (file) {
                                    uploadFile(file, {}, () => {}).then(
                                      (resp: any) => {
                                        let files = selectedAction?.files ?? [];
                                        if (
                                          files.filter(
                                            (f) =>
                                              !f.mime_type.includes("image")
                                          ).length === 0
                                        ) {
                                          files = [...files, resp.data];
                                        } else {
                                          files = files.map((f) => {
                                            if (
                                              !f.mime_type.includes("image")
                                            ) {
                                              return resp.data;
                                            }
                                            return f;
                                          });
                                        }
                                        setSelectedAction({
                                          ...selectedAction,
                                          files: files,
                                        });
                                      }
                                    );
                                  }
                                }
                              );
                            }
                          }}
                        />
                      </div>
                      <BsTrash
                        size={20}
                        className="absolute bottom-2 right-2 cursor-pointer text-red-400 hover:text-red-600"
                        onClick={() => {
                          if (
                            (selectedAction?.files ?? []).find(
                              (f) => !f.mime_type.includes("image")
                            )
                          ) {
                            const confirmFirst = window.confirm(
                              "Are you sure you want to delete this file?"
                            );
                            if (!confirmFirst) return;

                            deleteFile(
                              (selectedAction?.files ?? []).find(
                                (f) => !f.mime_type.includes("image")
                              )?.id!
                            ).then(() => {
                              setSelectedAction({
                                ...selectedAction!,
                                files: (selectedAction?.files ?? []).filter(
                                  (f) => f.mime_type.includes("image")
                                ),
                              });
                            });
                          }
                        }}
                      />
                    </div>
                  </div>
                </div> */}
              </>
            )}
            {selectedAction?.action_trigger === "IDLE" && (
              <>
                <div>
                  <Label htmlFor="idle_time" value="Idle Time" />
                  <TextInput
                    type="number"
                    value={selectedAction?.action_data?.idle_time}
                    onChange={(e) =>
                      setSelectedAction({
                        ...selectedAction!,
                        action_data: {
                          ...selectedAction!.action_data!,
                          idle_time: Number(e.target.value),
                        },
                      })
                    }
                    placeholder="Idle Time"
                  />
                </div>
                <div>
                  <Label htmlFor="idle_time_type" value="Idle Time Periode" />
                  <Select
                    options={idleTimeType}
                    value={idleTimeType.find(
                      (e) =>
                        e.value === selectedAction?.action_data!.idle_time_type
                    )}
                    onChange={(e) =>
                      setSelectedAction({
                        ...selectedAction!,
                        action_data: {
                          ...selectedAction!.action_data!,
                          idle_time_type: e!.value!,
                        },
                      })
                    }
                  />
                </div>
                <div>
                  <Label htmlFor="hour" value="Hour" />
                  <div className="flex col-span-2">
                    <input
                      type="time"
                      id="time"
                      className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                      value={selectedAction?.action_hour}
                      onChange={(e) => {
                        setSelectedAction({
                          ...selectedAction!,
                          action_data: {
                            ...selectedAction!.action_data!,
                          },
                          action_hour: e!.target.value,
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
                </div>
                <div>
                  <Label htmlFor="run_once" value="Run Once" />
                  <div className="flex col-span-2">
                    <ToggleSwitch
                      onChange={(e) =>
                        setSelectedAction({
                          ...selectedAction!,
                          run_once: e,
                        })
                      }
                      checked={selectedAction?.run_once ?? false}
                    />
                  </div>
                </div>
                {/*                 
                <div>
                  <ToggleSwitch
                    onChange={(e) =>
                      setSelectedAction({
                        ...selectedAction!,
                        status: e ? "ACTIVE" : "INACTIVE",
                      })
                    }
                    checked={selectedAction?.status === "ACTIVE"}
                  />
                </div> */}
              </>
            )}
          </div>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button
            onClick={() => {
              editColumnAction(projectId!, column!.id!, selectedAction!).then(
                () => {
                  setSelectedAction(undefined);
                  getColumn(projectId!, column!.id!).then((resp: any) => {
                    // console.log(resp.data)
                    setActions(resp.data.actions);
                    // onChangeColumn(resp.data);
                    // onAddAction(resp.data);
                  });
                }
              );
            }}
          >
            Save
          </Button>
        </Modal.Footer>
      </Modal>
    </>
  );
};
export default ModalColumn;

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
