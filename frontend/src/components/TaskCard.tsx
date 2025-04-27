import { useEffect, type FC } from "react";
import { TaskModel } from "../models/task";
import { Avatar } from "flowbite-react";
import { colorToStyle, getColor, initial, money } from "../utils/helper";
import { GoCommentDiscussion } from "react-icons/go";
import { BsCheck2Circle } from "react-icons/bs";
import { FormFieldType } from "../models/form";
import Moment from "react-moment";
import { Link } from "react-router-dom";

interface TaskCardProps {
  task: TaskModel;
  isdragged: boolean;
  onClick: (task: TaskModel) => void;
}


const TaskCard: FC<TaskCardProps> = ({ task, isdragged = false, onClick }) => {
  useEffect(() => {}, [task]);

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
        case FormFieldType.Price:
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
    
  return (
    <div
      className={`bg-white p-2 mt-2 rounded shadow`}
      id={task?.id}
      style={{
        transform: isdragged ? "rotate(-2deg)" : undefined,
        borderLeftColor: colorToStyle(getColor(task?.percentage ?? 0)),
        borderLeftWidth: "4px",
      }}
      onClick={() => onClick(task)}
    >
      <div className="flex items-center justify-between">
        {task?.name}
        <div className="flex gap-2 flex-row items-center">
          {task.completed && (
            <BsCheck2Circle className="text-green-500" size={16} />
          )}
          {(task?.comment_count ?? 0) > 0 && (
            <div className="text-xs flex-row flex gap-1 items-center">
              {(task?.comment_count ?? 0) < 100
                ? task?.comment_count ?? 0
                : "+99"}
              <GoCommentDiscussion />
            </div>
          )}
          <Avatar
            rounded
            img={task?.assignee?.user?.profile_picture?.url}
            alt="Avatar"
            size="xs"
            placeholderInitials={initial(task?.assignee?.user?.full_name)}
          />
        </div>
      </div>
      {task.task_attribute && (
        <div>
          <ul className="text-xs">
            {(task.task_attribute?.fields??[])
              .filter((e) => e.is_pinned)
              .map((e) => (
                <li key={e.id}>{renderValue(e.type, e.value)}</li>
              ))}
          </ul>
        </div>
      )}
    </div>
  );
};
export default TaskCard;
