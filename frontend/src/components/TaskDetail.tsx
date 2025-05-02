import { Editor } from "@tinymce/tinymce-react";
import {
  Avatar,
  Button,
  Carousel,
  Checkbox,
  Datepicker,
  FileInput,
  Label,
  Modal,
  Popover,
  Progress,
  Radio,
  Select as SelectFlowBite,
  Tabs,
  TabsRef,
  Textarea,
  TextInput,
  ToggleSwitch,
  Tooltip,
} from "flowbite-react";
import moment from "moment";
import {
  ReactNode,
  useContext,
  useEffect,
  useRef,
  useState,
  type FC,
} from "react";
import toast from "react-hot-toast";
import { AiOutlineLike } from "react-icons/ai";
import {
  BsActivity,
  BsAsterisk,
  BsCheck2Circle,
  BsCollection,
  BsFloppy,
  BsPencil,
  BsQuote,
  BsTrash,
  BsWhatsapp,
  BsYoutube,
} from "react-icons/bs";
import { FaDigg, FaRetweet } from "react-icons/fa6";
import { GoComment, GoCommentDiscussion } from "react-icons/go";
import { HiRefresh } from "react-icons/hi";
import { IoCheckmarkDone, IoShareOutline } from "react-icons/io5";
import { LuLink } from "react-icons/lu";
import { PiBookmark, PiPlay, PiPlayCircle } from "react-icons/pi";
import { RiAttachment2, RiFullscreenFill } from "react-icons/ri";
import { SiGoogleforms } from "react-icons/si";
import Markdown from "react-markdown";
import { Mention, MentionsInput } from "react-mentions";
import Moment from "react-moment";
import { Link } from "react-router-dom";
import Select, { InputActionMeta } from "react-select";
import remarkGfm from "remark-gfm";
import { ActiveCompanyContext } from "../contexts/CompanyContext";
import { LoadingContext } from "../contexts/LoadingContext";
import { ProfileContext } from "../contexts/ProfileContext";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { FormField, FormFieldType } from "../models/form";
import { ProjectModel, ProjectPreference } from "../models/project";
import {
  CompanyRapidApiPluginModel,
  RapidApiDataModel,
  RapidApiEndpointModel,
} from "../models/rapid_api";
import { TaskCommentModel, TaskModel } from "../models/task";
import { TaskAttributeModel } from "../models/task_attribute";
import { WhatsappMessageModel } from "../models/whatsapp_message";
import { getCompanyRapidAPIPlugins } from "../services/api/commonApi";
import { generateContent } from "../services/api/geminiApi";
import {
  addComment,
  addTaskPlugin,
  deleteTask,
  getTask,
  getTaskPluginData,
  getTaskPlugins,
  updateTask,
} from "../services/api/taskApi";
import { getTaskAttributes } from "../services/api/taskAttributeApi";
import {
  createWAMessage,
  getWhatsappMessages,
  markAsRead,
} from "../services/api/whatsappApi";
import { priorityOptions, severityOptions } from "../utils/constants";
import { getColor, initial, invert, money, nl2br } from "../utils/helper";
import { parseMentions } from "../utils/helper-ui";

interface TaskDetailProps {
  task: TaskModel;
  project: ProjectModel;
  onSwitchFullscreen: () => void;
}

const TaskDetail: FC<TaskDetailProps> = ({
  task,
  project,
  onSwitchFullscreen,
}) => {
  const { activeCompany, setActiveCompany } = useContext(ActiveCompanyContext);
  const [preference, setPreference] = useState<ProjectPreference>();
  const { loading, setLoading } = useContext(LoadingContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const tabsRef = useRef<TabsRef>(null);
  const [messages, setMessages] = useState<WhatsappMessageModel[]>([]);

  const [activeTab, setActiveTab] = useState(0);
  const [mounted, setMounted] = useState(false);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [comment, setComment] = useState("");
  const [isEditted, setIsEditted] = useState(false);
  const [activeTask, setActiveTask] = useState<TaskModel>();
  const [editDesc, setEditDesc] = useState(false);
  const [watchers, setWatchers] = useState<
    { label: string; value: string; avatar: ReactNode }[]
  >([]);
  const [comments, setComments] = useState<TaskCommentModel[]>([]);
  const [emojis, setEmojis] = useState([]);
  const [addPlugin, setAddPlugin] = useState(false);
  const [companyPlugins, setCompanyPlugins] = useState<
    CompanyRapidApiPluginModel[]
  >([]);
  const [selectedPlugin, setSelectedPlugin] =
    useState<CompanyRapidApiPluginModel>();
  const [selectedEnpoint, setSelectedEnpoint] =
    useState<RapidApiEndpointModel>();
  const [pluginTitle, setPluginTitle] = useState("");
  const [pluginEndpointParams, setPluginEndpointParams] = useState<
    { key: string; type: string; value: string }[]
  >([]);
  const [pluginDatas, setPluginDatas] = useState<RapidApiDataModel[]>([]);
  const [modalAi, setModalAi] = useState(false);
  const [aiPrompt, setAiPrompt] = useState("");
  const [taskAttributes, setTaskAttributes] = useState<TaskAttributeModel[]>(
    []
  );
  const chatContainerRef = useRef<HTMLDivElement>(null);
  const timeout = useRef<number | null>(null);
  const [content, setContent] = useState("");
  const [openAttachment, setOpenAttachment] = useState(false);
  useEffect(() => {
    if (!activeTask?.ref_id) return;
    if (
      wsMsg?.session_id == activeTask?.ref_id &&
      wsMsg?.command == "WHATSAPP_RECEIVED"
    ) {
      setMessages([...messages, wsMsg.data]);
      setTimeout(() => {
        scrollToBottom();
      }, 300);
    }
  }, [wsMsg, profile, activeTask?.ref_id]);

  const scrollToBottom = () => {
    const element = document.getElementById("channel-messages");
    if (element) {
      element.scrollTo({
        top: element.scrollHeight,
        behavior: "smooth",
      });
    }
  };

  useEffect(() => {
    scrollToBottom();
  }, [messages]);

  const handleScroll = () => {
    const messageElements = document.querySelectorAll(".message");

    messageElements.forEach((el) => {
      const observer = new IntersectionObserver(
        (entries) => {
          entries.forEach((entry) => {
            if (entry.isIntersecting) {
              // const messageId = parseInt(entry.target.dataset.id);
              // markAsRead(messageId);
              let message = messages.find(
                (m) => m.id == entry.target.getAttribute("id")
              );
              // console.log(message?.message)
              if (message && !message.is_read) {
                message.is_read = true;
                setMessages([
                  ...messages.map((m) => {
                    if (m.id == message?.id) {
                      return { ...m, is_read: true };
                    }
                    return m;
                  }),
                ]);

                timeout.current = window.setTimeout(() => {
                  markAsRead(message!.id!);
                }, 500);
              }
            }
          });
        },
        { threshold: 0.3 } // Minimal 50% pesan terlihat
      );

      observer.observe(el);
    });
  };
  useEffect(() => {
    fetch(process.env.REACT_APP_BASE_URL + "/assets/static/emojis.json")
      .then((response) => {
        return response.json();
      })
      .then((jsonData) => {
        setEmojis(jsonData.emojis);
      });
  }, []);
  const neverMatchingRegex = /($a)/;
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
    setMounted(true);
  }, []);
  useEffect(() => {
    if (task != activeTask) {
      // console.log(task != activeTask)
      //   setIsEditted(true);
    }
    if (mounted && preference?.rapid_api_enabled) {
      getAllCompanyPlugins();
      if (project && activeTask) {
        getTaskPlugins(project!.id!, activeTask!.id!).then((resp: any) => {
          setPluginDatas(resp.data);
        });
      }
    }
    if (mounted && preference?.custom_attribute_enabled) {
      getTaskAttributes({ page: 1, size: 20 }).then((v: any) =>
        setTaskAttributes(v.data.items)
      );
    }
  }, [mounted, preference]);

  useEffect(() => {
    setPluginEndpointParams(JSON.parse(selectedEnpoint?.params ?? "[]"));
  }, [selectedEnpoint]);

  const getAllCompanyPlugins = async () => {
    try {
      // setLoading(true);
      const resp: any = await getCompanyRapidAPIPlugins();
      setCompanyPlugins(resp.data);
    } catch (error: any) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (!mounted) return;
    if (!activeTask) {
      getDetail(task.id!);
    }
  }, [task, mounted]);

  useEffect(() => {
    if (wsMsg?.project_id == task.project_id) {
      //   console.log("wsMsg", wsMsg);
      if (wsMsg.task_id == task.id) {
        getDetail(task.id!);
      }
    }
  }, [wsMsg]);

  const handleEditorChange = (e: any) => {
    setActiveTask({
      ...activeTask,
      description: e.target.getContent(),
      comments: [],
    });
    setIsEditted(true);
  };

  const getDetail = (id: string) => {
    getTask(task.project_id!, id)
      .then((resp: any) => {
        setActiveTask(resp.data);
        setPreference(resp.preference);
        if (resp.data?.ref_type == "whatsapp_session") {
          getWhatsappMessages(resp.data?.ref_id, {
            page: 1,
            size: 100,
            search: "",
          })
            .then((res: any) => {
              setMessages(res.data.items);
            })
            .catch((err) => {
              console.error(err);
              window.location.href = "/whatsapp";
            });
        }
      })
      .catch(toast.error);
  };

  const emojiStyle = {
    control: {
      fontSize: 16,
      lineHeight: 1.2,
      minHeight: 63,
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

  const saveTask = () => {
    updateTask(task!.project_id!, task.id!, {
      ...activeTask,
      watchers: watchers.map((watcher) => ({
        id: watcher.value,
      })),
    })
      .catch(toast.error)
      .then(() => {
        toast.success("Task updated successfully");
        setIsEditted(false);
      });
  };

  const renderValue = (fieldType: FormFieldType, val: any) => {
    switch (fieldType) {
      case FormFieldType.DateRangePicker:
        return (
          val && (
            <div>
              {val[0] && <Moment format="DD MMM YYYY">{val[0]}</Moment>} ~{" "}
              {val[1] && <Moment format="DD MMM YYYY">{val[1]}</Moment>}
            </div>
          )
        );
      case FormFieldType.DatePicker:
        return val && <Moment format="DD MMM YYYY">{val}</Moment>;
      case FormFieldType.PasswordField:
        return val && "* * * * * * *";
      case FormFieldType.ToggleSwitch:
        return val && <BsCheck2Circle />;
      case FormFieldType.FileUpload:
        return (
          val && (
            <Link to={val} target="_blank">
              {val}
            </Link>
          )
        );
      case FormFieldType.NumberField:
      case FormFieldType.Currency:
        return money(parseFloat(val));
      case FormFieldType.Checkbox:
        return (
          <ul>
            {val.map((e: any) => (
              <li key={e}>{e}</li>
            ))}
          </ul>
        );

      default:
        break;
    }
    return val;
  };

  useEffect(() => {
    if (activeTask) {
      setWatchers(
        (activeTask?.watchers ?? []).map((m: any) => ({
          label: m?.user?.full_name,
          value: m.id,
          avatar: (
            <Avatar
              rounded
              img={m?.user?.profile_picture?.url}
              alt="Avatar"
              size="xs"
              color="blue"
              placeholderInitials={initial(m?.user?.full_name)}
            />
          ),
        }))
      );
    }
  }, [activeTask]);

  const renderCommentBox = () => (
    <div className="flex flex-row gap-4 items-start p-2">
      <Avatar
        rounded
        img={profile?.profile_picture?.url}
        placeholderInitials={initial(profile?.full_name)}
        alt="Avatar"
        size="sm"
      />
      <div className="flex-1">
        {/* <h3 className="text-xl font-semibold text-gray-600 mb-2">
          {profile?.full_name}
        </h3> */}
        <MentionsInput
          value={comment}
          onChange={(val) => {
            setComment(val.target.value);
          }}
          style={emojiStyle}
          placeholder={
            "Press ':' for emojis, mention people using '@' and shift+enter to send"
          }
          autoFocus
          onKeyDown={async (val: any) => {
            if (val.key === "Enter" && val.shiftKey) {
              try {
                setLoading(true);
                if (comment) {
                  await addComment(project!.id!, task.id!, {
                    comment,
                  });

                  setComment("");

                  toast.success("Comment added successfully");
                  // tabsRef.current?.setActiveTab(1);
                }
              } catch (error) {
                toast.error(`${error}`);
              } finally {
                setLoading(false);
              }

              return;
            }
          }}
        >
          <Mention
            trigger="@"
            data={(project?.members ?? []).map((member) => ({
              id: member.id!,
              display: member.user?.full_name!,
            }))}
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
        {/* <div className="flex justify-end mt-4">
        <Button
          className="w-32"
          onClick={() => {
            if (comment) {
              addComment(project!.id!, task.id!, {
                comment,
              })
                .catch(toast.error)
                .then(() => {
                  setComment("");

                  toast.success("Comment added successfully");
                  tabsRef.current?.setActiveTab(1);
                });
            }
          }}
        >
          Submit
        </Button>
      </div> */}
      </div>
    </div>
  );

  const renderAttributeField = (field: FormField) => {
    switch (field.type) {
      case FormFieldType.TextField:
        return (
          <div key={field.id}>
            <TextInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e.target.value;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e.target.value };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.PasswordField:
        return (
          <div key={field.id}>
            <TextInput
              type="password"
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e.target.value;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e.target.value };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.FileUpload:
        return (
          <div key={field.id}>
            <FileInput
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              onChange={(e) => {
                const file = e.target.files;
                if (file) {
                  const reader = new FileReader();
                  reader.onload = (event) => {
                    field.value = event.target?.result as string;

                    setActiveTask({
                      ...activeTask,
                      task_attribute: {
                        ...activeTask!.task_attribute!,
                        fields:
                          activeTask!.task_attribute?.fields.map((f) => {
                            if (f.id === field.id) {
                              return { ...f, value: e.target.value };
                            }
                            return f;
                          }) ?? [],
                      },
                    });
                  };
                  reader.readAsDataURL(file[0]);
                }

                // field?.value = e.target.value
                // setSectionValues(sectionValues)
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.EmailField:
        return (
          <div key={field.id}>
            <TextInput
              type={"email"}
              sizing={"sm"}
              name={field.label}
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e.target.value;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e.target.value };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.NumberField:
      case FormFieldType.Currency:
      case FormFieldType.Price:
        return (
          <div key={field.id}>
            <TextInput
              type={"number"}
              sizing={"sm"}
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e.target.value;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e.target.value };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );

      case FormFieldType.TextArea:
        return (
          <div key={field.id}>
            <Textarea
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e.target.value;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e.target.value };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DatePicker:
        return (
          <div key={field.id}>
            <Datepicker
              placeholder={field.placeholder}
              required={field.required}
              value={field?.value}
              onChange={(e) => {
                field.value = e;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return { ...f, value: e };
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.DateRangePicker:
        if (!field?.value) {
          field.value = [new Date(), new Date()];
        }
        return (
          <div className="">
            <div className="grid grid-cols-2 gap-2">
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
                value={moment(field?.value[0]).toDate()}
                onChange={(e) => {
                  // let val = [...field?.value];
                  setActiveTask({
                    ...activeTask,
                    task_attribute: {
                      ...activeTask!.task_attribute!,
                      fields:
                        activeTask!.task_attribute?.fields.map((f) => {
                          if (f.id === field.id) {
                            return { ...f, value: [e, field.value[1]] };
                          }
                          return f;
                        }) ?? [],
                    },
                  });
                }}
              />
              <Datepicker
                placeholder={field.placeholder}
                required={field.required}
                value={moment(field?.value[1]).toDate()}
                onChange={(e) => {
                  setActiveTask({
                    ...activeTask,
                    task_attribute: {
                      ...activeTask!.task_attribute!,
                      fields:
                        activeTask!.task_attribute?.fields.map((f) => {
                          if (f.id === field.id) {
                            return { ...f, value: [field.value[0], e] };
                          }
                          return f;
                        }) ?? [],
                    },
                  });
                }}
              />
            </div>
            {field.help_text && (
              <small className="text-gray-400">{field.help_text}</small>
            )}
          </div>
        );
      case FormFieldType.RadioButton:
        return (
          <div key={field.id}>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}
              {field.options.map((option, i) => (
                <div className="flex items-center gap-2" key={i}>
                  <Radio
                    id={`${field.label}-${i}`}
                    value={option.value}
                    checked={field?.value == option.value}
                    onChange={(val) => {
                      // field?.value =
                      //   val.target.value;

                      setActiveTask({
                        ...activeTask,
                        task_attribute: {
                          ...activeTask!.task_attribute!,
                          fields:
                            activeTask!.task_attribute?.fields.map((f) => {
                              if (f.id === field.id) {
                                return { ...f, value: val.target.value };
                              }
                              return f;
                            }) ?? [],
                        },
                      });
                    }}
                  />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.Checkbox:
        if (!field?.value) {
          field.value = [];
        }
        return (
          <div key={field.id}>
            <fieldset className="flex max-w-md flex-col gap-4">
              {field.help_text && (
                <legend className="mb-4 text-sm text-gray-400">
                  {field.help_text}
                </legend>
              )}

              {field.options.map((option, i) => (
                <div className="flex items-center gap-2" key={i}>
                  <Checkbox
                    id={`${field.label}-${i}`}
                    value={option.value}
                    checked={field.value.includes(option.value)}
                    onChange={(val) => {
                      if (!field.value.includes(option.value)) {
                        field.value.push(option.value);
                      } else {
                        field.value = [
                          ...field.value.filter(
                            (val: any) => val != option.value
                          ),
                        ];
                      }

                      setActiveTask({
                        ...activeTask,
                        task_attribute: {
                          ...activeTask!.task_attribute!,
                          fields:
                            activeTask!.task_attribute?.fields.map((f) => {
                              if (f.id === field.id) {
                                return field;
                              }
                              return f;
                            }) ?? [],
                        },
                      });
                    }}
                  />
                  <Label htmlFor={option.value}>{option.label}</Label>
                </div>
              ))}
            </fieldset>
          </div>
        );
      case FormFieldType.ToggleSwitch:
        if (!field?.value) {
          field.value = false;
        }
        return (
          <div key={field.id}>
            <ToggleSwitch
              sizing="sm"
              checked={field?.value}
              label={field.help_text}
              onChange={(val) => {
                field.value = val;
                setActiveTask({
                  ...activeTask,
                  task_attribute: {
                    ...activeTask!.task_attribute!,
                    fields:
                      activeTask!.task_attribute?.fields.map((f) => {
                        if (f.id === field.id) {
                          return field;
                        }
                        return f;
                      }) ?? [],
                  },
                });
              }}
            />
          </div>
        );
      case FormFieldType.Dropdown:
        return (
          <div key={field.id}>
            <SelectFlowBite
              value={field?.value}
              onChange={(val) => {
                field.value = val.target.value;
              }}
            >
              {field.options.map((option, i) => (
                <option
                  selected={field?.value == option.value}
                  key={i}
                  value={option.value}
                >
                  {option.label}
                </option>
              ))}
            </SelectFlowBite>
          </div>
        );
    }
  };

  return (
    <div className="flex flex-col h-full w-full">
      <div className="flex-1 space-y-2 overflow-y-auto">
        <Progress
          color={getColor(activeTask?.percentage ?? 0)}
          size="sm"
          progress={activeTask?.percentage ?? 0}
        />
        <div className="flex flex-row items-center justify-between">
          <input
            className="border-0 py-2 text-2xl font-semibold focus:border-0 focus:outline-none w-full"
            value={activeTask?.name ?? ""}
            onChange={(el) => {
              setIsEditted(true);
              setActiveTask({ ...activeTask, name: el.target.value });
            }}
          />
          <div className="flex flex-row gap-2 items-center">
            <Tooltip content={`Save`} placement="left">
              <BsFloppy
                className="text-gray-400 hover:text-gray-600 cursor-pointer"
                onClick={() => {
                  updateTask(task!.project_id!, task.id!, {
                    ...activeTask,
                    watchers: watchers.map((watcher) => ({
                      id: watcher.value,
                    })),
                    task_attribute_data: activeTask?.task_attribute
                      ? JSON.stringify(activeTask?.task_attribute)
                      : "{}",
                  })
                    .catch(toast.error)
                    .then(() => {
                      toast.success("Task updated successfully");
                      setIsEditted(false);
                    });
                }}
              />
            </Tooltip>
            {!activeTask?.completed && (
              <div className="relative">
                <TextInput
                  type="number"
                  max={100}
                  value={activeTask?.percentage ?? 0}
                  onChange={(e) => {
                    setIsEditted(true);
                    setActiveTask({
                      ...activeTask,
                      percentage:
                        parseInt(e.target.value) > 100
                          ? 100
                          : parseInt(e.target.value),
                    });
                  }}
                  className="w-20 p-0 !text-right"
                  sizing="sm"
                  style={{ textAlign: "right", paddingRight: 20 }}
                />
                <small className="absolute top-1/2 -translate-y-1/2 right-2">
                  %
                </small>
              </div>
            )}

            {activeTask?.completed ? (
              <Tooltip
                content={`Completed at ${moment(
                  activeTask?.completed_date
                ).format("DD MMM YYYY, HH:mm")}`}
                placement="left"
              >
                <BsCheck2Circle className="text-green-500" />
              </Tooltip>
            ) : (
              <Button
                size="xs"
                onClick={() => {
                  updateTask(task!.project_id!, task.id!, {
                    ...activeTask,
                    completed: true,
                  })
                    .catch(toast.error)
                    .then(() => {
                      toast.success("Task updated successfully");
                      setIsEditted(false);
                      setActiveTask({ ...activeTask, completed: true });
                    });
                }}
                color="gray"
                className="w-40"
              >
                <BsCheck2Circle className="mr-2" />
                Mark As Completed
              </Button>
            )}

            <Tooltip content="Full Screen" placement="left">
              <RiFullscreenFill
                className="cursor-pointer"
                onClick={onSwitchFullscreen}
              />
            </Tooltip>
            <Tooltip content="Delete Task" placement="left">
              <BsTrash
                className="cursor-pointer text-red-400 hover:text-red-600"
                onClick={() => {
                  if (
                    window.confirm(
                      "Are you sure you want to delete this task? This action is irreversible."
                    )
                  ) {
                    deleteTask(task.id!)
                      .catch(toast.error)
                      .then(() => {
                        toast.success("Task deleted successfully");
                        setIsEditted(false);
                        setActiveTask(undefined);
                      });
                  }
                }}
              />
            </Tooltip>
          </div>
        </div>
        <table className="w-full">
          <tbody>
            <tr>
              <td className="px-2 py-1 w-28"> Date</td>
              <td className="px-2 py-1 w-28">
                <Datepicker
                  className="!border-0"
                  value={moment(activeTask?.start_date).toDate()}
                  onChange={(date) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, start_date: date! });
                  }}
                />
              </td>
              <td className="px-2 py-1 w-28"> Assigned To</td>
              <td className="px-2 py-1 w-28">
                <select
                  className="w-full rounded-lg border-gray-200 bg-gray-50"
                  value={activeTask?.assignee_id!}
                  onChange={(el) => {
                    setIsEditted(true);
                    setActiveTask({
                      ...activeTask,
                      assignee_id: el.target.value,
                    });
                  }}
                >
                  <option value={""}>Select</option>
                  {(project?.members ?? []).map((member) => (
                    <option key={member.id} value={member.id}>
                      {member.user?.full_name}
                    </option>
                  ))}
                </select>
              </td>
            </tr>
            <tr>
              <td className="px-2 py-1 w-28"> Due Date</td>
              <td className="px-2 py-1 w-28">
                <Datepicker
                  className="!border-0"
                  value={moment(activeTask?.end_date).toDate()}
                  onChange={(date) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, end_date: date! });
                  }}
                />
              </td>
              <td className="px-2 py-1 w-28"> Watcher</td>
              <td className="px-2 py-1 w-28">
                <Select
                  className="w-full"
                  styles={{
                    multiValue: (styles, { data }) => {
                      return {
                        ...styles,
                        backgroundColor: "#f0f0f0",
                        borderRadius: "5px",
                      };
                    },
                  }}
                  isMulti
                  value={watchers}
                  onChange={(values) => {
                    setIsEditted(true);
                    setWatchers(values.map((val) => val));
                  }}
                  options={(project?.members ?? []).map((member) => ({
                    label: member.user?.full_name ?? "",
                    value: member.id ?? "",
                    avatar: (
                      <Avatar
                        rounded
                        img={member?.user?.profile_picture?.url}
                        alt="Avatar"
                        size="xs"
                        placeholderInitials={initial(member?.user?.full_name)}
                        color="blue"
                      />
                    ),
                  }))}
                  formatOptionLabel={(option: any) => (
                    <div className="flex flex-row gap-2  items-center">
                      {option.avatar}
                      <span>{option.label}</span>
                    </div>
                  )}
                  inputValue={""}
                  onInputChange={(
                    newValue: string,
                    actionMeta: InputActionMeta
                  ) => {
                    // console.log(newValue, actionMeta);
                  }}
                />
              </td>
            </tr>
            <tr>
              <td className="px-2 py-1 w-28"> Priority</td>
              <td className="px-2 py-1 w-28">
                <Select
                  className="w-full"
                  isSearchable={false}
                  defaultValue={priorityOptions.find(
                    (option) => option.value === task?.priority
                  )}
                  onChange={(val) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, priority: val?.value! });
                  }}
                  options={priorityOptions}
                  formatOptionLabel={(option: any) => (
                    <div
                      className="flex flex-row gap-2  items-center text-center px-2 py-1 rounded"
                      style={{ backgroundColor: option.color }}
                    >
                      <span style={{ color: invert(option.color) }}>
                        {option.label}
                      </span>
                    </div>
                  )}
                  inputValue={""}
                  onInputChange={(
                    newValue: string,
                    actionMeta: InputActionMeta
                  ) => {
                    // console.log(newValue, actionMeta);
                  }}
                />
              </td>
              <td className="px-2 py-1 w-28"> Severity</td>
              <td className="px-2 py-1 w-28">
                <Select
                  className="w-full"
                  isSearchable={false}
                  defaultValue={severityOptions.find(
                    (option) => option.value === task?.severity
                  )}
                  onChange={(val) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, severity: val?.value! });
                  }}
                  options={severityOptions}
                  formatOptionLabel={(option: any) => (
                    <div
                      className="flex flex-row gap-2  items-center text-center px-2 py-1 rounded"
                      style={{ backgroundColor: option.color }}
                    >
                      <span style={{ color: invert(option.color) }}>
                        {option.label}
                      </span>
                    </div>
                  )}
                  inputValue={""}
                  onInputChange={(
                    newValue: string,
                    actionMeta: InputActionMeta
                  ) => {
                    // console.log(newValue, actionMeta);
                  }}
                />
              </td>
            </tr>
          </tbody>
        </table>
        {editDesc ? (
          <Editor
            apiKey={process.env.REACT_APP_TINY_MCE_KEY}
            init={{
              plugins:
                "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount ",
              toolbar:
                "closeButton saveButton aiButton | undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat ",
              setup: (editor: any) => {
                editor.ui.registry.addButton("closeButton", {
                  icon: "close",
                  tooltip: "Close editor",
                  onAction: (_: any) => setEditDesc(false),
                });

                editor.ui.registry.addButton("saveButton", {
                  icon: "save",
                  tooltip: "Save Task",
                  onAction: (_: any) => {
                    saveTask();
                  },
                });
                editor.ui.registry.addButton("aiButton", {
                  icon: "ai",
                  tooltip: "Ai",
                  onAction: (_: any) => {
                    if (activeCompany?.setting?.gemini_api_key) {
                      setModalAi(true);
                    } else {
                      toast.error("Please add gemini api key");
                    }
                  },
                });

                editor.ui.registry.addMenuItem("closeButton", {
                  text: "Close editor",
                  onAction: (_: any) => setEditDesc(false),
                });

                editor.ui.registry.addMenuItem("saveButton", {
                  text: "Save",
                  onAction: (_: any) => {
                    saveTask();
                  },
                });
              },
              menubar: "file edit view insert format tools table custom",
              menu: {
                custom: { title: "Editor", items: "closeButton saveButton" },
              },
            }}
            initialValue={activeTask?.description ?? ""}
            onChange={handleEditorChange}
          />
        ) : (
          <div>
            <div className="flex flex-row gap-2 items-center group/item">
              <h3 className="font-semibold text-xl">Description</h3>
              <BsPencil
                size={16}
                className="cursor-pointer group/edit invisible group-hover/item:visible text-gray-600"
                onClick={() => setEditDesc((prev) => !prev)}
              />
            </div>
            <div
              className="min-h-10"
              dangerouslySetInnerHTML={{
                __html: activeTask?.description ?? "",
              }}
            />
          </div>
        )}
        {companyPlugins.length > 0 && preference?.rapid_api_enabled && (
          <div>
            <div className="flex justify-between items-center">
              <h3 className="font-semibold text-xl">Plugins</h3>
              <Button
                onClick={() => setAddPlugin(true)}
                className="mt-4"
                color="gray"
                size="xs"
              >
                Add Plugin
              </Button>
            </div>
            {pluginDatas.map((e) => (
              <div key={e.id}>
                <div className="flex gap-2 items-center">
                  <h3 className="text-lg font-semibold">{e.title}</h3>
                  <HiRefresh
                    className="text-gray-400 hover:text-gray-600 cursor-pointer"
                    onClick={() => {
                      getTaskPluginData(project!.id!, activeTask!.id!, e.id)
                        .catch((err) => toast.error(`${err}`))
                        .then(() => {
                          toast.success("Plugin refresh successfully");
                          getTaskPlugins(project!.id!, activeTask!.id!).then(
                            (resp: any) => {
                              setPluginDatas(resp.data);
                            }
                          );
                        });
                    }}
                  />
                </div>
                <div className="flex flex-row gap-2">
                  {e.parsed_data?.thumbnails && (
                    <div className="aspect-square w-[300px] object-cover">
                      <Carousel leftControl="<" rightControl=">">
                        {(e.parsed_data?.thumbnails ?? []).map(
                          (thumbail: any, i: number) => (
                            <img src={thumbail} key={i} className=" " />
                          )
                        )}
                      </Carousel>
                    </div>
                  )}
                  <div className="w-full">
                    <div className="flex justify-between">
                      <h3 className="text-xl font[500] mb-2">Insight</h3>
                    </div>
                    <table className="w-full">
                      <tbody>
                        {e.parsed_params[0].value && (
                          <tr>
                            <td className="pr-2">Link</td>
                            <td className="text-right"></td>
                            <td className="w-1">
                              <LuLink
                                size={12}
                                className=" text-blue-400 hover:text-blue-600 cursor-pointer"
                                onClick={() => {
                                  if (
                                    e.parsed_data?.endpoint_key ==
                                    "video_details"
                                  ) {
                                    window.open(
                                      `https://www.youtube.com/watch?v=${e.parsed_data?.id}`
                                    );
                                    // window.open(e.parsed_params[0].value);
                                  }
                                  if (
                                    e.parsed_data?.endpoint_key ==
                                    "get_tweet_details"
                                  ) {
                                    window.open(
                                      `https://x.com/${e.parsed_data?.user?.screen_name}/status/${e.parsed_data?.id_str}`
                                    );
                                    // window.open(e.parsed_params[0].value);
                                  }
                                  if (
                                    e.parsed_data?.endpoint_key ==
                                    "get_facebook_post_details"
                                  ) {
                                    window.open(e.parsed_data?.post_link);
                                    // window.open(e.parsed_params[0].value);
                                  }
                                  if (
                                    e.parsed_data?.endpoint_key ==
                                    "get_post_detail"
                                  ) {
                                    window.open(
                                      `https://www.tiktok.com/@${e.parsed_data?.author.uniqueId}/video/${e.parsed_data?.id}?is_from_webapp=1&sender_device=pc`
                                    );
                                    // window.open(e.parsed_params[0].value);
                                  }

                                  if (
                                    e.parsed_data?.endpoint_key ==
                                    "media_info_by_url_v2_media_info_by_url_get"
                                  ) {
                                    window.open(e.parsed_params[0].value);
                                  }
                                }}
                              />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions && (
                          <tr>
                            <td className="pr-2">Reactions</td>
                            <td className="text-right"></td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Angry && (
                          <tr>
                            <td className="pr-2">Angry</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Angry)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Care && (
                          <tr>
                            <td className="pr-2">Care</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Care)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Haha && (
                          <tr>
                            <td className="pr-2">Haha</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Haha)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Love && (
                          <tr>
                            <td className="pr-2">Love</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Love)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Sad && (
                          <tr>
                            <td className="pr-2">Sad</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Sad)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.Wow && (
                          <tr>
                            <td className="pr-2">Wow</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reactions?.Wow)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.reactions?.total_reaction_count && (
                          <tr>
                            <td className="pr-2">Total Reaction Count</td>
                            <td className="text-right">
                              {money(
                                e.parsed_data?.reactions?.total_reaction_count
                              )}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.comments_count && (
                          <tr>
                            <td className="pr-2">Comments</td>
                            <td className="text-right">
                              {money(e.parsed_data?.comments_count)}
                            </td>
                            <td className="w-1">
                              <GoComment size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.viewCount && (
                          <tr>
                            <td className="pr-2">Views</td>
                            <td className="text-right">
                              {money(e.parsed_data?.viewCount)}
                            </td>
                            <td className="w-1">
                              <BsYoutube size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.likeCount && (
                          <tr>
                            <td className="pr-2">Likes</td>
                            <td className="text-right">
                              {money(e.parsed_data?.likeCount)}
                            </td>
                            <td className="w-1">
                              <AiOutlineLike size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.commentCount && (
                          <tr>
                            <td className="pr-2">Comments</td>
                            <td className="text-right">
                              {money(e.parsed_data?.commentCount)}
                            </td>
                            <td className="w-1">
                              <GoComment size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.share_count && (
                          <tr>
                            <td className="pr-2">Share Count</td>
                            <td className="text-right">
                              {money(e.parsed_data?.share_count)}
                            </td>
                            <td className="w-1"></td>
                          </tr>
                        )}
                        {e.parsed_data?.stats?.collectCount && (
                          <tr>
                            <td className="pr-2">Collects</td>
                            <td className="text-right">
                              {money(e.parsed_data?.stats?.collectCount)}
                            </td>
                            <td className="w-1">
                              <BsCollection size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.stats?.commentCount && (
                          <tr>
                            <td className="pr-2">Comments</td>
                            <td className="text-right">
                              {money(e.parsed_data?.stats?.commentCount)}
                            </td>
                            <td className="w-1">
                              <GoComment size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.stats?.playCount && (
                          <tr>
                            <td className="pr-2">Plays</td>
                            <td className="text-right">
                              {money(e.parsed_data?.stats?.playCount)}
                            </td>
                            <td className="w-1">
                              <PiPlayCircle size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.stats?.shareCount && (
                          <tr>
                            <td className="pr-2">Shares</td>
                            <td className="text-right">
                              {money(e.parsed_data?.stats?.shareCount)}
                            </td>
                            <td className="w-1">
                              <IoShareOutline size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.stats?.diggCount && (
                          <tr>
                            <td className="pr-2">Diggs</td>
                            <td className="text-right">
                              {money(e.parsed_data?.stats?.diggCount)}
                            </td>
                            <td className="w-1">
                              <FaDigg size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.favorite_count && (
                          <tr>
                            <td className="pr-2">Favorites</td>
                            <td className="text-right">
                              {money(e.parsed_data?.favorite_count)}
                            </td>
                            <td className="w-1">
                              <AiOutlineLike size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.like_count && (
                          <tr>
                            <td className="pr-2">Likes</td>
                            <td className="text-right">
                              {money(e.parsed_data?.like_count)}
                            </td>
                            <td className="w-1">
                              <AiOutlineLike size={12} />
                            </td>
                          </tr>
                        )}

                        {e.parsed_data?.reply_count && (
                          <tr>
                            <td className="pr-2">Replies</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reply_count)}
                            </td>
                            <td className="w-1">
                              <GoComment size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.comment_count && (
                          <tr>
                            <td className="pr-2">Comments</td>
                            <td className="text-right">
                              {money(e.parsed_data?.comment_count)}
                            </td>
                            <td className="w-1">
                              <GoComment size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.play_count && (
                          <tr>
                            <td className="pr-2">Plays</td>
                            <td className="text-right">
                              {money(e.parsed_data?.play_count)}
                            </td>
                            <td className="w-1">
                              <PiPlay size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.retweet_count && (
                          <tr>
                            <td className="pr-2">Retweets</td>
                            <td className="text-right">
                              {money(e.parsed_data?.retweet_count)}
                            </td>
                            <td className="w-1">
                              <FaRetweet size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.quote_count && (
                          <tr>
                            <td className="pr-2">Quotes</td>
                            <td className="text-right">
                              {money(e.parsed_data?.quote_count)}
                            </td>
                            <td className="w-1">
                              <BsQuote size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.save_count && (
                          <tr>
                            <td className="pr-2">Saves</td>
                            <td className="text-right">
                              {money(e.parsed_data?.save_count)}
                            </td>
                            <td className="w-1">
                              <PiBookmark size={12} />
                            </td>
                          </tr>
                        )}
                        {e.parsed_data?.reshare_count && (
                          <tr>
                            <td className="pr-2">Reshares</td>
                            <td className="text-right">
                              {money(e.parsed_data?.reshare_count)}
                            </td>
                            <td className="w-1">
                              <IoShareOutline size={12} />
                            </td>
                          </tr>
                        )}
                      </tbody>
                    </table>

                    {e.parsed_data?.caption && (
                      <div
                        className="text-xs mt-4"
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.caption.text),
                        }}
                      />
                    )}
                    {e.parsed_data?.full_text && (
                      <div
                        className="text-xs mt-4"
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.full_text),
                        }}
                      />
                    )}
                    {e.parsed_data?.desc && (
                      <div
                        className="text-xs mt-4"
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.desc),
                        }}
                      />
                    )}
                    {e.parsed_data?.values?.text && (
                      <div
                        className="text-xs mt-4"
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.values?.text),
                        }}
                      />
                    )}
                    {e.parsed_data?.title && (
                      <div
                        className="text-sm font-[500] mt-4 "
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.title),
                        }}
                      />
                    )}
                    {e.parsed_data?.description && (
                      <div
                        className="text-xs mt-4"
                        dangerouslySetInnerHTML={{
                          __html: nl2br(e.parsed_data?.description),
                        }}
                      />
                    )}
                  </div>
                </div>
              </div>
            ))}
          </div>
        )}
        <Tabs
          aria-label="Default tabs"
          variant="default"
          ref={tabsRef}
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
          className="mt-4"
        >
          {/* <Tabs.Item active title="Add Comments" icon={GoComment}>
            {renderCommentBox()}
          </Tabs.Item> */}
          <Tabs.Item
            title={
              <div className="flex flex-row gap-1">
                {(task?.comment_count ?? 0) > 0 && (
                  <div className=" flex-row flex gap-2 items-center">
                    {(task?.comment_count ?? 0) < 100
                      ? task?.comment_count ?? 0
                      : "+99"}
                  </div>
                )}
                Comments
              </div>
            }
            icon={GoCommentDiscussion}
          >
            <div className=" ">
              <div className="h-[calc(50vh-60px)]   overflow-y-auto">
                <div className="space-y-2">
                  {(activeTask?.comments ?? []).map((comment) => (
                    <div
                      key={comment.id}
                      className={`flex items-start gap-2 p-2  dark:bg-gray-800 ${
                        profile?.id == comment?.member?.user_id
                          ? "justify-end"
                          : "justify-start"
                      }`}
                    >
                      {profile?.id != comment?.member?.user_id && (
                        <Avatar
                          rounded
                          img={comment.member?.user?.profile_picture?.url}
                          alt="Avatar"
                          size="sm"
                          placeholderInitials={initial(
                            comment.member?.user?.full_name
                          )}
                          className="py-2"
                        />
                      )}

                      <div className="flex flex-col rounded bg-gray-100 p-2 min-w-[300px] ">
                        <div className="flex justify-between items-end">
                          <span className="font-medium text-gray-800 dark:text-white">
                            {comment.member?.user?.full_name}
                          </span>
                          <Moment fromNow className="text-xs">
                            {comment?.published_at}
                          </Moment>
                        </div>
                        <span className="text-gray-600 dark:text-gray-300">
                          {parseMentions(comment.comment ?? "", (type, id) => {
                            // console.log(type, id);
                          })}
                        </span>
                      </div>
                    </div>
                  ))}
                  {/* <div className="flex justify-center">
                  <Button
                    color="gray"
                    onClick={() => {
                      tabsRef.current?.setActiveTab(0);
                    }}
                  >
                    Add New Comment
                  </Button>
                </div> */}
                </div>
              </div>
              <div className="h-[60px]">{renderCommentBox()}</div>
            </div>
          </Tabs.Item>
          <Tabs.Item title="Activity" icon={BsActivity}>
            <div className="h-[calc(50vh-60px)]   overflow-y-auto">
              <ul className="space-y-4">
                {activeTask?.activities?.map((activity) => (
                  <li
                    key={activity.id}
                    className="p-2 bg-gray-100 px-4 py-2 w-fit rounded-lg"
                  >
                    <span className="text-gray-600 dark:text-gray-300 hover:font-semibold">
                      {activity.member?.user?.full_name}
                    </span>{" "}
                    <strong>
                      {activity.activity_type?.replaceAll("_", " ")}
                    </strong>{" "}
                    at <Moment fromNow>{activity.activity_date}</Moment>
                  </li>
                ))}
              </ul>
            </div>
          </Tabs.Item>
          {activeTask?.form_response && preference?.form_enabled && (
            <Tabs.Item title="Form Response" icon={SiGoogleforms}>
              {(activeTask?.form_response?.sections ?? []).map((e) => (
                <div className="" key={e.id}>
                  <h2 className="text-lg font-bold">{e.section_title}</h2>

                  <table className="w-full mb-4">
                    {e.fields.map((f) => (
                      <tr key={f.id} className="border">
                        <td className="px-2 py-1 font-semibold bg-gray-50 border w-[300px]">
                          {f.label}
                        </td>

                        <td className="px-2 py-1 border">
                          {renderValue(f.type, f.value)}
                        </td>
                      </tr>
                    ))}
                  </table>
                </div>
              ))}
            </Tabs.Item>
          )}
          {preference?.custom_attribute_enabled && (
            <Tabs.Item title="Attributes" icon={BsAsterisk}>
              <table className="w-full">
                <tr>
                  <td className="px-2 py-1 w-1/3"> Attribute</td>
                  <td className="px-2 py-1">
                    <Select
                      className="w-full"
                      isSearchable={false}
                      defaultValue={{
                        label: activeTask?.task_attribute?.title,
                        value: activeTask?.task_attribute_id,
                      }}
                      // value={activeTask?.task_attribute_id}
                      onChange={(val) => {
                        setIsEditted(true);
                        setActiveTask({
                          ...activeTask,
                          task_attribute_id: val?.value,
                        });
                      }}
                      options={taskAttributes.map((e) => ({
                        label: e.title,
                        value: e.id,
                      }))}
                      inputValue={""}
                      onInputChange={(
                        newValue: string,
                        actionMeta: InputActionMeta
                      ) => {
                        // console.log(newValue, actionMeta);
                      }}
                    />
                  </td>
                </tr>
                {(activeTask?.task_attribute?.fields ?? []).map((e) => (
                  <tr>
                    <td className="px-2 py-1 w-1/3"> {e.label}</td>
                    <td className="px-2 py-1">{renderAttributeField(e)} </td>
                  </tr>
                ))}
              </table>
            </Tabs.Item>
          )}
          {task?.ref_type == "whatsapp_session" && (
            <Tabs.Item title="WhatsApp" icon={BsWhatsapp}>
              <div
                id="channel-messages"
                className="messages h-[calc(50vh-120px)] overflow-y-auto p-4 bg-gray-50 space-y-8"
                ref={chatContainerRef}
                onScroll={handleScroll}
              >
                {messages.map((msg) => (
                  <div
                    key={msg.id}
                    className={`message flex flex-row items-end mb-2  ${
                      msg.is_from_me ? "justify-end" : "justify-start"
                    }`}
                    id={msg.id}
                  >
                    <div
                      className={`min-w-[300px] max-w-[600px] ${
                        !msg.is_from_me
                          ? "bg-green-500 text-white"
                          : "bg-gray-200"
                      } p-2 rounded-md`}
                      data-id={msg.id}
                    >
                      {msg.media_url && msg.mime_type?.includes("video") && (
                        <video
                          controls
                          src={msg.media_url}
                          className={`rounded-md mb-2 ${
                            msg.is_from_me ? "ml-auto" : "mr-auto"
                          } w-[300px] h-[300px] object-cover`}
                        />
                      )}
                      {msg.media_url && msg.mime_type?.includes("audio") && (
                        <audio
                          controls
                          src={msg.media_url}
                          className={`rounded-md mb-2 ${
                            msg.is_from_me ? "ml-auto" : "mr-auto"
                          } w-[300px]`}
                        />
                      )}

                      {msg.media_url && msg.mime_type?.includes("image") && (
                        <Popover
                          placement="bottom"
                          content={
                            <div className="bg-white p-4 rounded-md w-[600px]">
                              <img
                                src={msg.media_url}
                                alt=""
                                className="w-full h-full object-cover rounded-md"
                              />
                            </div>
                          }
                        >
                          <img
                            src={msg.media_url}
                            alt=""
                            className={` rounded-md mb-2 ${
                              msg.is_from_me ? "ml-auto" : "mr-auto"
                            } w-[300px] h-[300px] object-cover`}
                          />
                        </Popover>
                      )}
                      {!msg.is_from_me && <small>{msg.contact?.name}</small>}
                      {msg.is_group && !msg.is_from_me && (
                        <small>{msg.message_info?.PushName}</small>
                      )}

                      <Markdown remarkPlugins={[remarkGfm]}>
                        {msg.message}
                      </Markdown>
                      <div className="text-[10px] justify-between flex items-center">
                        {msg.sent_at && <Moment fromNow>{msg.sent_at}</Moment>}
                        {msg.is_read && (
                          <IoCheckmarkDone
                            size={16}
                            style={{
                              color: msg.is_from_me ? "green" : "white",
                            }}
                          />
                        )}
                      </div>
                    </div>
                  </div>
                ))}
              </div>
              <div className="shoutbox border-t pt-2 min-h-[20px] max-h[60px] px-2  flex justify-between items-center gap-2">
                <MentionsInput
                  value={content}
                  onChange={(val: any) => {
                    setContent(val.target.value);
                  }}
                  style={emojiStyle}
                  placeholder={"Press ':' for emojis and shift+enter to send"}
                  className="w-full"
                  autoFocus
                  onKeyDown={async (val: any) => {
                    if (val.key === "Enter" && val.shiftKey) {
                      setContent((prev) => prev + "\n");
                      return;
                    }
                    if (val.key === "Enter") {
                      try {
                        await createWAMessage(task!.ref_id!, {
                          message: content,
                          // files: files,
                        });
                        // setOpenAttachment(false);
                        // setFiles([]);
                      } catch (error) {
                        toast.error(`${error}`);
                      } finally {
                        setTimeout(() => {
                          setContent("");
                        }, 300);
                      }

                      return;
                    }
                  }}
                >
                  <Mention
                    trigger=":"
                    markup="__id__"
                    regex={neverMatchingRegex}
                    data={queryEmojis}
                  />
                </MentionsInput>
                <Button color="gray" onClick={() => setOpenAttachment(true)}>
                  <RiAttachment2 />
                </Button>
              </div>
            </Tabs.Item>
          )}
        </Tabs>
      </div>
      {/* {isEditted && (
        <div className="bg-red border-t pt-2 flex flex-row justify-end gap-2">
          <Button
            className="w-32"
            color="gray"
            onClick={() => {
              updateTask(task!.project_id!, task.id!, {
                ...activeTask,
                watchers: watchers.map((watcher) => ({
                  id: watcher.value,
                })),
              })
                .catch(toast.error)
                .then(() => {
                  toast.success("Task updated successfully");
                  setIsEditted(false);
                });
            }}
          >
            Save
          </Button>
        </div>
      )} */}
      <Modal show={addPlugin} onClose={() => setAddPlugin(false)}>
        <Modal.Header>Add Plugin</Modal.Header>
        <Modal.Body>
          <div className="space-y-4 flex flex-col">
            <div className="flex flex-col gap-1 ">
              <Label htmlFor="plugin-name">Title</Label>
              <TextInput
                placeholder="Title"
                value={pluginTitle}
                onChange={(el) => {
                  setPluginTitle(el.target.value);
                }}
              />
            </div>
            <div className="flex flex-col gap-1 ">
              <Label htmlFor="plugin-name">Plugin Name</Label>
              <Select
                id="plugin-name"
                value={
                  selectedPlugin
                    ? {
                        label: selectedPlugin.rapid_api_plugin?.name,
                        value: selectedPlugin.rapid_api_plugin_id,
                      }
                    : null
                }
                onChange={(option) => {
                  setSelectedPlugin(
                    companyPlugins.find(
                      (e) => e.rapid_api_plugin_id == option?.value
                    )
                  );
                  setSelectedEnpoint(undefined);
                }}
                options={companyPlugins.map((plugin) => ({
                  label: plugin.rapid_api_plugin?.name,
                  value: plugin.rapid_api_plugin_id,
                }))}
                placeholder="Select a plugin"
              />
            </div>
            {selectedPlugin && (
              <div className="flex flex-col gap-1 ">
                <Label htmlFor="plugin-name">End Point</Label>
                <Select
                  id="plugin-name"
                  value={
                    selectedEnpoint
                      ? {
                          label: selectedEnpoint.title,
                          value: selectedEnpoint.id,
                        }
                      : null
                  }
                  onChange={(option) => {
                    setSelectedEnpoint(
                      (
                        selectedPlugin?.rapid_api_plugin?.rapid_api_endpoints ??
                        []
                      ).find((e) => e.id == option?.value)
                    );
                  }}
                  options={(
                    selectedPlugin?.rapid_api_plugin?.rapid_api_endpoints ?? []
                  ).map((endpoint) => ({
                    label: endpoint.title,
                    value: endpoint.id,
                  }))}
                  placeholder="Select a plugin"
                />
              </div>
            )}
            {selectedPlugin && selectedEnpoint && (
              <div>
                {pluginEndpointParams.map((e: any, i: number) => (
                  <div className="flex flex-col gap-1 " key={i}>
                    <Label htmlFor="plugin-name">{e["key"]}</Label>
                    <TextInput
                      value={e["value"]}
                      placeholder={e["key"]}
                      onChange={(el) => {
                        setPluginEndpointParams([
                          ...pluginEndpointParams.map((p, j) => {
                            if (j == i) {
                              return { ...p, value: el.target.value };
                            }
                            return p;
                          }),
                        ]);
                      }}
                    />
                  </div>
                ))}
              </div>
            )}
            <div className="h-[160px]"></div>
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex justify-end">
            <Button
              onClick={() => {
                addTaskPlugin(project!.id!, activeTask!.id!, {
                  task_id: activeTask!.id,
                  title: pluginTitle,
                  rapid_api_endpoint_id: selectedEnpoint!.id,
                  rapid_api_plugin_id: selectedPlugin!.rapid_api_plugin_id,
                  params: JSON.stringify(pluginEndpointParams),
                })
                  .then((v) => {
                    toast.success("Plugin added successfully");
                    setAddPlugin(false);
                    getTaskPlugins(project!.id!, activeTask!.id!).then(
                      (resp: any) => {
                        setPluginDatas(resp.data);
                      }
                    );
                  })
                  .catch(toast.error);
              }}
            >
              Save
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
      <Modal show={modalAi} onClose={() => setModalAi(false)}>
        <Modal.Header>AI Prompt</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col gap-4">
            <Label htmlFor="ai-prompt">Enter AI Prompt</Label>
            <Textarea
              id="ai-prompt"
              value={aiPrompt}
              onChange={(e) => setAiPrompt(e.target.value)}
              placeholder="Type your AI prompt here"
            />
          </div>
        </Modal.Body>
        <Modal.Footer>
          <div className="flex justify-end">
            <Button
              onClick={() => {
                // Handle the AI prompt submission logic here
                // setModalAi(false);

                // toast.success("AI prompt submitted successfully");
                let prompt = `${aiPrompt} 
pastikan format responsenya :
{
    "description": "ini deskripsi",
}
            
                `;
                generateContent(prompt, "", true, true).then((resp: any) => {
                  if (resp.data.description) {
                    // toast.success(resp.data.description)
                    setActiveTask({
                      ...activeTask,
                      description: resp.data.description,
                    });
                    setModalAi(false);
                    toast.success("AI prompt submitted successfully");
                  }
                });
              }}
            >
              Submit
            </Button>
          </div>
        </Modal.Footer>
      </Modal>
    </div>
  );
};
export default TaskDetail;
