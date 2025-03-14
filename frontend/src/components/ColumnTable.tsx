import type { FC } from "react";
import { ColumnModel } from "../models/column";
import { TaskModel } from "../models/task";
import Moment from "react-moment";
import { Avatar } from "flowbite-react";
import { initial } from "../utils/helper";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import TaskTable from "./TaskTable";

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

  return (
    <div className="bg-white p-4 rounded-lg" {...attributes} ref={setNodeRef} style={{ backgroundColor: column.color }}>
      <div className="flex gap-2">
        <div >{column.icon}</div>
        <h3 className="text-lg font-semibold" {...listeners}>
          {column.name}
        </h3>
      </div>
      <table className="w-full ">
        <thead>
          <tr className="bg-gray-50">
            <td className="px-2 py-2 border w-1/4 font-semibold">Task Name</td>
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
              <TaskTable key={index} task={task} onSelectTask={onSelectTask} />
            ))}
            {column.tasks?.length === 0 && (
              <TaskTable task={null} onSelectTask={(val) => {}} />
            )}
          </tbody>
        </SortableContext>
      </table>
    </div>
  );
};
export default ColumnTable;
