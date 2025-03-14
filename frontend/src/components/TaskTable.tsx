import type { FC } from "react";
import { TaskModel } from "../models/task";
import Moment from "react-moment";
import { Avatar } from "flowbite-react";
import { initial } from "../utils/helper";
import { useSortable } from "@dnd-kit/sortable";
import { UniqueIdentifier } from "@dnd-kit/core";
import { randomUUID } from "crypto";

interface TaskTableProps {
  task: TaskModel | null;
  onSelectTask: (task: TaskModel) => void;
}

const TaskTable: FC<TaskTableProps> = ({ task, onSelectTask }) => {
  const { attributes, listeners, setNodeRef, transform, transition } =
    useSortable({
      id: task ? (task?.id as UniqueIdentifier) : crypto.randomUUID(),
    });

  if (!task)
    return (
      <tr  ref={setNodeRef}
      {...attributes}
      >
        <td colSpan={4} className="text-center py-4 text-gray-500 border">
          No tasks available.
        </td>
      </tr>
    );
  return (
    <tr
      className="bg-white select-none"
      ref={setNodeRef}
      {...attributes}
      {...listeners}
    >
      <td
        className="px-2 py-2 border min-w-[200px] hover:font-semibold"
        onClick={() => task && onSelectTask(task)}
      >
        {task?.name}
      </td>
      <td className="px-2 py-2 border min-w-[200px]">
        <Moment format="DD MMM YYYY">{task?.start_date}</Moment>
        {task?.end_date && (
          <span>
            {" ~ "} <Moment format="DD MMM YYYY">{task?.end_date}</Moment>
          </span>
        )}
      </td>
      <td className="px-2 py-2 border min-w-[200px]">
        {task?.assignee && (
          <div className="flex gap-2 items-center">
            <Avatar
              rounded
              img={task?.assignee.user?.profile_picture?.url}
              size="xs"
              placeholderInitials={initial(task?.assignee?.user?.full_name)}
            />
            {task?.assignee?.user?.full_name}
          </div>
        )}
      </td>
      <td className="px-2 py-2 border min-w-[200px]">
        <Avatar.Group>
          {task?.watchers?.map((member) => (
            <Avatar
              key={member?.user?.id}
              size="xs"
              img={member?.user?.profile_picture?.url}
              rounded
              stacked
              placeholderInitials={initial(member?.user?.full_name)}
            />
          ))}
          {(task?.watchers ?? []).length > 5 && (
            <Avatar.Counter
              total={(task?.watchers ?? []).length - 5}
              href="#"
            />
          )}
        </Avatar.Group>
      </td>
    </tr>
  );
};
export default TaskTable;
