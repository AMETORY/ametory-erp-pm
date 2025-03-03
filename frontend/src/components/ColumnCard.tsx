import { useContext, useEffect, useState, type FC } from "react";
import { ColumnModel } from "../models/column";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import { Droppable } from "./droppable";
import { RiDragMoveFill } from "react-icons/ri";
import { TaskModel } from "../models/task";
import { createTask, getTasks } from "../services/api/taskApi";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";

interface ColumnCardProps {
  projectId: string;
  column: ColumnModel;
  columns: ColumnModel[];
  onChange: (columns: ColumnModel[]) => void;
  onSelectTask: (task: TaskModel) => void;
}

const ColumnCard: FC<ColumnCardProps> = ({
  projectId,
  column,
  columns,
  onChange,
  onSelectTask,
}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [tasks, setTasks] = useState<TaskModel[]>([]);
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

  const getAllTasks = () => {
    getTasks(projectId, column.id as string).then((resp: any) => {
      setTasks(resp.data.items);
      column.tasks = resp.data.items;
    });
  }
  useEffect(() => {
    getAllTasks()
  }, []);

  useEffect(() => {
    if (!wsMsg) return
    if (!column.id) return
    // if (profile?.id != null && profile?.id == wsMsg?.sender_id) {
    if (wsMsg.column_id == column.id || wsMsg.source_column_id == column.id) {
      getAllTasks()
    }
    // }
  }, [wsMsg, profile]);

  return (
    <div
      {...attributes}
      ref={setNodeRef}
      className="bg-gray-200  p-4  rounded-lg inline-block overflw-y-scroll"
      key={column.id}
      style={{
        backgroundColor: column.color,
        width: "300px",
        minHeight: "calc(100vh - 180px)",
      }}
    >
      <div className="flex flex-row justify-between group/item">
        <h2 className="text-xl font-bold mb-2 text-gray-600 hover:text-black">
          {column.icon && <span className="text-sm mr-2"> {column.icon}</span>}{" "}
          {column.name}
        </h2>
        <RiDragMoveFill {...listeners} className="" />
      </div>
      <SortableContext
        id={column.id!}
        items={(column.tasks ?? []).map((item) => ({
          id: item.id as UniqueIdentifier,
        }))}
        strategy={verticalListSortingStrategy}
      >
        <Droppable
          column={column}
          onSelectTask={onSelectTask}
          onSelectItem={(id: string) => {
            console.log(id);
          }}
          onAddItem={() => {
            let totalItem = 0;
            for (const element of columns) {
              totalItem += (element.tasks ?? []).length;
            }
            (column.tasks ?? []).push({
              id: `item-${totalItem + 1}`,
              name: `Item ${totalItem + 1}`,
            });
            // console.log("add item",[...columns]);
            createTask(projectId, {
              name: `Item ${totalItem + 1}`,
              column_id: column.id,
            });
            // onChange([...columns]);
          }}
        />
      </SortableContext>
    </div>
  );
};
export default ColumnCard;
