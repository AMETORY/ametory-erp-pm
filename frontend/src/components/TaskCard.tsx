import { useEffect, type FC } from "react";
import { TaskModel } from "../models/task";
import { Avatar } from "flowbite-react";
import { initial } from "../utils/helper";

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
      <div className="flex items-center justify-between">
        {task?.name}
        <Avatar rounded img={task?.assignee?.user?.profile_picture?.url} alt="Avatar" size="xs" placeholderInitials={initial(task?.assignee?.user?.full_name)}/>
      </div>

    </div>
  );
};
export default TaskCard;
