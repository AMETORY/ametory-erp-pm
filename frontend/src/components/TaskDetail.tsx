import { Editor } from "@tinymce/tinymce-react";
import {
  Avatar,
  Button,
  Datepicker,
  Tabs,
  TabsRef,
  Tooltip,
} from "flowbite-react";
import {
  ReactNode,
  useContext,
  useEffect,
  useRef,
  useState,
  type FC,
} from "react";
import toast from "react-hot-toast";
import { HiAdjustments, HiClipboardList, HiUserCircle } from "react-icons/hi";
import { MdComment, MdDashboard } from "react-icons/md";
import { RiFullscreenFill } from "react-icons/ri";
import Select, { InputActionMeta } from "react-select";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProjectModel } from "../models/project";
import { TaskCommentModel, TaskModel } from "../models/task";
import { addComment, getTask, updateTask } from "../services/api/taskApi";
import { initial } from "../utils/helper";
import { BsActivity } from "react-icons/bs";
import { GoComment, GoCommentDiscussion } from "react-icons/go";
import { ProfileContext } from "../contexts/ProfileContext";
import { MentionsInput, Mention } from "react-mentions";
import { parseMentions } from "../utils/helper-ui";
import Moment from "react-moment";

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
  const { profile, setProfile } = useContext(ProfileContext);
  const tabsRef = useRef<TabsRef>(null);

  const [activeTab, setActiveTab] = useState(0);
  const [mounted, setMounted] = useState(false);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [comment, setComment] = useState("");
  const [isEditted, setIsEditted] = useState(false);
  const [activeTask, setActiveTask] = useState<TaskModel>();
  const [watchers, setWatchers] = useState<
    { label: string; value: string; avatar: ReactNode }[]
  >([]);
  const [comments, setComments] = useState<TaskCommentModel[]>([]);
  const [emojis, setEmojis] = useState([]);

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
  }, [activeTask]);

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
      .then((resp: any) => setActiveTask(resp.data))
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
  return (
    <div className="flex flex-col h-full w-full">
      <div className="flex-1 space-y-2 overflow-y-auto">
        <div className="flex flex-row items-center justify-between">
          <input
            className="border-0 py-2 text-2xl font-semibold focus:border-0 focus:outline-none w-full"
            value={activeTask?.name ?? ""}
            onChange={(el) => {
              setIsEditted(true);
              setActiveTask({ ...activeTask, name: el.target.value });
            }}
          />
          <Tooltip content="Full Screen" placement="left">
            <RiFullscreenFill
              className="cursor-pointer"
              onClick={onSwitchFullscreen}
            />
          </Tooltip>
        </div>
        <table className="w-full">
          <tbody>
            <tr>
              <td className="px-2 py-1 w-28"> Date</td>
              <td className="px-2 py-1 w-28">
                <Datepicker
                  className="!border-0"
                  value={activeTask?.start_date!}
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
              <td className="px-2 py-1 w-28"> Date Line</td>
              <td className="px-2 py-1 w-28">
                <Datepicker
                  className="!border-0"
                  value={activeTask?.end_date!}
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
          </tbody>
        </table>
        <Editor
          apiKey={process.env.REACT_APP_TINY_MCE_KEY}
          init={{
            plugins:
              "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount ",
            toolbar:
              "undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat",
          }}
          initialValue={activeTask?.description ?? ""}
          onChange={handleEditorChange}
        />
        <Tabs
          aria-label="Default tabs"
          variant="default"
          ref={tabsRef}
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
        >
          <Tabs.Item active title="Add Comments" icon={GoComment}>
            <div className="flex flex-row gap-4 items-start">
              <Avatar
                rounded
                img={profile?.profile_picture?.url}
                placeholderInitials={initial(profile?.full_name)}
                alt="Avatar"
                size="md"
              />
              <div className="flex-1">
                <h3 className="text-xl font-semibold text-gray-600 mb-2">
                  {profile?.full_name}
                </h3>
                {/* <Editor
                  apiKey={process.env.REACT_APP_TINY_MCE_KEY}
                  init={{
                    plugins:
                      "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount textcolor colorpicker",
                    toolbar: "emoticons | blocks fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor",
                    menubar: false,
                    statusbar: false,
                    height: 150,
                  }}
                  onChange={(e: any) => {
                    setComment(e.target.getContent());
                  }}
                /> */}
                <MentionsInput
                  value={comment}
                  onChange={(val) => {
                    setComment(val.target.value);
                  }}
                  style={emojiStyle}
                  placeholder={"Press ':' for emojis, mention people using '@'"}
                  autoFocus
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
                <div className="flex justify-end mt-4">
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
                </div>
              </div>
            </div>
          </Tabs.Item>
          <Tabs.Item title="Comments" icon={GoCommentDiscussion}>
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
              <div className="flex justify-center">
                <Button
                  color="gray"
                  onClick={() => {
                    tabsRef.current?.setActiveTab(0);
                  }}
                >
                  Add New Comment
                </Button>
              </div>
            </div>
          </Tabs.Item>
          <Tabs.Item title="Activity" icon={BsActivity}>
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
          </Tabs.Item>
        </Tabs>
        
      </div>
      {isEditted && (
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
      )}
    </div>
  );
};
export default TaskDetail;
