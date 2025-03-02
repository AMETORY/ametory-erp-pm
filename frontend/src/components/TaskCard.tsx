import { useEffect, type FC } from "react";
import { TaskModel } from "../models/task";

interface TaskCardProps {
  task: TaskModel;
  isdragged: boolean;
  onClick: (task: TaskModel) => void
}

const TaskCard: FC<TaskCardProps> = ({ task, isdragged = false, onClick }) => {
  useEffect(() => {}, [task]);
  return (
    <div
      className="bg-white p-2 mt-2 rounded shadow"
      id={task?.id}
      style={{
        transform: isdragged ? "rotate(-2deg)" : undefined,
      }}
      onClick={() => onClick(task)}
    >
      {task?.name}
    </div>
  );
};
export default TaskCard;
