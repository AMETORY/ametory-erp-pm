import { useContext, useEffect, useState, type FC } from "react";
import { ColumnModel } from "../models/column";
import { TaskModel } from "../models/task";
import Moment from "react-moment";
import { Avatar, Button } from "flowbite-react";
import { initial } from "../utils/helper";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import TaskTable from "./TaskTable";
import { BsPencil } from "react-icons/bs";
import ModalColumn from "./ModalColumn";
import { createTask, getTasks } from "../services/api/taskApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import { SearchContext } from "../contexts/SearchContext";

interface ColumnTableProps {
  projectId: string;
  column: ColumnModel;
  columns: ColumnModel[];
  onChange: (columns: ColumnModel[]) => void;
  onChangeColumn: (column: ColumnModel) => void;
  onSelectTask: (task: TaskModel) => void;
  onAddItem: (taskId: string) => void;
}

const ColumnTable: FC<ColumnTableProps> = ({
  projectId,
  column,
  columns,
  onChange,
  onChangeColumn,
  onSelectTask,
  onAddItem,
}) => {
  const {search, setSearch} = useContext(SearchContext);
  const {
    isOver,
    setNodeRef,
    setActivatorNodeRef,
    attributes,
    listeners,
    transform,
    transition,
  } = useSortable({
    id: column.id as UniqueIdentifier,
  });
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [tasks, setTasks] = useState<TaskModel[]>([]);

  const getAllTasks = () => {
    getTasks(projectId, column.id as string, search).then((resp: any) => {
      setTasks(resp.data.items);
      column.tasks = resp.data.items;
      onChangeColumn(column);
    });
  };
  useEffect(() => {
    if (!wsMsg) return;
    if (!column.id) return;
    // if (profile?.id != null && profile?.id == wsMsg?.sender_id) {
    if (wsMsg.column_id == column.id || wsMsg.source_column_id == column.id) {
      getAllTasks();
    }
    // }
  }, [wsMsg, profile,search]);

  const [showModal, setShowModal] = useState(false);
  return (
    <div
      className="bg-white p-4 rounded-lg"
      {...attributes}
      ref={setNodeRef}
      style={{ backgroundColor: column.color }}
    >
      <div className="flex justify-between group/item">
        <div className="flex gap-2">
          <div>{column.icon}</div>
          <h3 className="text-lg font-semibold" {...listeners}>
            {column.name}
          </h3>
        </div>
        <div className="group/edit invisible group-hover/item:visible flex gap-2 items-center">
          <Button
            size="xs"
            color="transparent"
            onClick={() => {
              setShowModal(true);
            }}
          >
            <BsPencil className=" text-gray-600 mr-2" /> Edit
          </Button>
          <Button
            size="xs"
            color="transparent"
            onClick={() => {
              let totalItem = 0;
              for (const element of columns) {
                totalItem += (element.tasks ?? []).length;
              }

              // console.log("add item",[...columns]);
              createTask(projectId, {
                name: `Task #${totalItem + 1}`,
                column_id: column.id,
                order_number: (column.tasks ?? []).length + 1,
              }).then((resp: any) => {
                onAddItem(resp.task_id);
              });
            }}
          >
            + Add Task
          </Button>
        </div>
      </div>
      <table className="w-full ">
        <thead>
          <tr className="bg-gray-50">
            <td
              className="px-2 py-2 border w-1/4 font-semibold"
              style={{
                borderLeftColor: column?.color,
                borderLeftWidth: "4px",
              }}
            >
              Task Name
            </td>
            <td className="px-2 py-2 border w-1/4 font-semibold">Date</td>
            <td className="px-2 py-2 border w-1/4 font-semibold">Assignee</td>
            <td className="px-2 py-2 border w-1/4 font-semibold">Watcher</td>
          </tr>
        </thead>
        <SortableContext
          id={column.id!}
          items={(column.tasks ?? []).map((item) => ({
            id: item.id as UniqueIdentifier,
          }))}
          strategy={verticalListSortingStrategy}
        >
          <tbody>
            {(column.tasks ?? []).map((task, index) => (
              <TaskTable
                column={column}
                key={index}
                task={task}
                onSelectTask={onSelectTask}
              />
            ))}
            {column.tasks?.length === 0 && (
              <TaskTable
                column={column}
                task={null}
                onSelectTask={(val) => {}}
              />
            )}
          </tbody>
        </SortableContext>
      </table>
      <ModalColumn
        showModal={showModal}
        setShowModal={setShowModal}
        projectId={projectId}
        onChangeColumn={onChangeColumn}
        column={column}
        onAddAction={() => {

        }}
      />
    </div>
  );
};
export default ColumnTable;


// http://localhost:8081/api/v1/project/a9e5c750-d0b9-49fd-a50c-03849dc4b232/column/1da9939f-5a3e-4c4d-b5d4-7a79f90d859c/add-action