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
import { TaskModel } from "../models/task";
import { getTask, updateTask } from "../services/api/taskApi";
import { initial } from "../utils/helper";
import { BsActivity } from "react-icons/bs";
import { GoComment, GoCommentDiscussion } from "react-icons/go";
import { ProfileContext } from "../contexts/ProfileContext";

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
    if (wsMsg?.project_id != task.project_id) {
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
    });
    setIsEditted(true);
  };

  const getDetail = (id: string) => {
    getTask(task.project_id!, id)
      .then((resp: any) => setActiveTask(resp.data))
      .catch(toast.error);
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
      <div className="flex-1 space-y-2">
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
              "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount textcolor colorpicker",
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
          onActiveTabChange={(tab) => setActiveTab(tab)}
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
                <h3 className="text-xl font-semibold text-gray-600 mb-2">{profile?.full_name}</h3>
                <Editor
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
                />
                <div className="flex justify-end mt-4">
                  <Button
                    className="w-32"
                    onClick={() => {
                      if (comment) {
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
            This is{" "}
            <span className="font-medium text-gray-800 dark:text-white">
              Profile tab's associated content
            </span>
            . Clicking another tab will toggle the visibility of this one for
            the next. The tab JavaScript swaps classes to control the content
            visibility and styling.
          </Tabs.Item>
          <Tabs.Item title="Activity" icon={BsActivity}>
            This is{" "}
            <span className="font-medium text-gray-800 dark:text-white">
              Dashboard tab's associated content
            </span>
            . Clicking another tab will toggle the visibility of this one for
            the next. The tab JavaScript swaps classes to control the content
            visibility and styling.
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
