import { useContext, useEffect, useState, type FC } from "react";
import { TaskModel } from "../models/task";
import { Button, Datepicker, TextInput } from "flowbite-react";
import { getTask, updateTask } from "../services/api/taskApi";
import { Editor } from "@tinymce/tinymce-react";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { debounce } from "../utils/helper";
import { ProjectModel } from "../models/project";

interface TaskDetailProps {
  task: TaskModel;
  project: ProjectModel;
}

const TaskDetail: FC<TaskDetailProps> = ({ task, project }) => {
  const [mounted, setMounted] = useState(false);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const [isEditted, setIsEditted] = useState(false);
  const [activeTask, setActiveTask] = useState<TaskModel>();

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
      .catch(console.error);
  };
  return (
    <div className="flex flex-col h-full w-full">
      <div className="flex-1 space-y-2">
        <div>
          <input
            className="border-0 py-2 text-2xl font-semibold focus:border-0 focus:outline-none w-full"
            value={activeTask?.name ?? ""}
            onChange={(el) => {
              setIsEditted(true);
              setActiveTask({ ...activeTask, name: el.target.value });
            }}
          />
        </div>
        <table className="w-full">
          <tbody>
            <tr>
              <td className="px-2 py-1 w-1/4"> Date</td>
              <td className="px-2 py-1 w-1/4">
                <Datepicker
                className="!border-0"
                  value={activeTask?.start_date!}
                  onChange={(date) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, start_date: date! });
                  }}
                />
              </td>
              <td className="px-2 py-1 w-1/4"> Assigned To</td>
              <td className="px-2 py-1 w-1/4">
                <select
                  className="w-full rounded-lg border-gray-200 bg-gray-50"
                  value={activeTask?.assignee_id!}
                  onChange={(el) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, assignee_id: el.target.value });
                  }}
                >
                  <option value={""}>Select</option>
                  {(project?.members??[]).map((member) => (
                    <option key={member.id} value={member.id}>{member.user?.full_name}</option>
                  ))}
                </select>
              </td>
            </tr>
            <tr>
          
              <td className="px-2 py-1 w-1/4"> Date Line</td>
              <td className="px-2 py-1 w-1/4">
                <Datepicker
                className="!border-0"
                  value={activeTask?.end_date!}
                  onChange={(date) => {
                    setIsEditted(true);
                    setActiveTask({ ...activeTask, end_date: date! });
                  }}
                />
              </td>
              <td className="px-2 py-1 w-1/4"> Watcher</td>
              <td className="px-2 py-1 w-1/4">
               
              </td>
            </tr>
          </tbody>
        </table>
        <Editor
          apiKey={process.env.REACT_APP_TINY_MCE_KEY}
          init={{
            plugins:
              "anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount",
            toolbar:
              "undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat",
          }}
          initialValue={activeTask?.description ?? ""}
          onChange={handleEditorChange}
        />
      </div>
      {isEditted && (
        <div className="bg-red border-t pt-2 flex flex-row justify-end gap-2">
          <Button
            className="w-32"
            color="gray"
            onClick={() => {
              updateTask(task!.project_id!, task.id!, activeTask).catch(
                console.error
              );
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
