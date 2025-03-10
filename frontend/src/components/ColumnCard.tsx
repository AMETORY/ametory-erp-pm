import { UniqueIdentifier } from "@dnd-kit/core";
import {
  SortableContext,
  useSortable,
  verticalListSortingStrategy,
} from "@dnd-kit/sortable";
import { Button, Label, Modal, TextInput } from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import toast from "react-hot-toast";
import { BsPencil } from "react-icons/bs";
import { RiDragMoveFill } from "react-icons/ri";
import { ProfileContext } from "../contexts/ProfileContext";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ColumnModel } from "../models/column";
import { TaskModel } from "../models/task";
import { updateColumn } from "../services/api/projectApi";
import { createTask, getTasks } from "../services/api/taskApi";
import { Droppable } from "./droppable";

interface ColumnCardProps {
  projectId: string;
  column: ColumnModel;
  columns: ColumnModel[];
  onChange: (columns: ColumnModel[]) => void;
  onChangeColumn: (column: ColumnModel) => void;
  onSelectTask: (task: TaskModel) => void;
  onAddItem: (taskId: string) => void;
}

const ColumnCard: FC<ColumnCardProps> = ({
  projectId,
  column,
  columns,
  onChange,
  onChangeColumn,
  onSelectTask,
  onAddItem,
}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [tasks, setTasks] = useState<TaskModel[]>([]);
  const [showModal, setShowModal] = useState(false);
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
      onChangeColumn(column);
    });
  };
  useEffect(() => {
    getAllTasks();
  }, []);

  useEffect(() => {
    if (!wsMsg) return;
    if (!column.id) return;
    // if (profile?.id != null && profile?.id == wsMsg?.sender_id) {
    if (wsMsg.column_id == column.id || wsMsg.source_column_id == column.id) {
      getAllTasks();
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
        minHeight: "calc(100vh - 230px)",
      }}
    >
      <div className="flex flex-row justify-between group/item">
        <div className="flex flex-1 items-center">
          {column.icon && (
            <h2 className="text-xl font-bold mb-2 text-gray-600 hover:text-black">
              <span className="text-sm mr-2"> {column.icon}</span>
            </h2>
          )}
          <input
            className="text-xl font-bold mb-2 text-gray-600 hover:text-black bg-transparent border-0 focus:ring-0 focus:outline-none rounded-lg m-0 p-0"
            type="text"
            value={column.name}
            onChange={(e) =>
              onChangeColumn({ ...column, name: e.target.value })
            }
            onKeyUp={(e) => {
              if (e.key === "Enter") {
                (e.target as HTMLInputElement).blur();
              }
            }}
            onBlur={(e) => {
              updateColumn(projectId!, {
                ...column,
              });
            }}
          />
        </div>
        <div className="flex gap-2">
          <BsPencil
            className="group/edit invisible group-hover/item:visible text-gray-600"
            onClick={() => setShowModal(true)}
          />
          <RiDragMoveFill
            {...listeners}
            className="group/edit invisible group-hover/item:visible"
          />
        </div>
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

            // console.log("add item",[...columns]);
            createTask(projectId, {
              name: `Task #${totalItem + 1}`,
              column_id: column.id,
              order_number: (column.tasks ?? []).length +1,
            }).then((resp: any) => {
              onAddItem(resp.task_id);
            })
            // onChange([...columns]);
          }}
        />
      </SortableContext>
      <Modal show={showModal} onClose={() => setShowModal(false)}>
        <Modal.Header>Edit Column</Modal.Header>
        <Modal.Body className="space-y-6">
          <div>
            <div className="mb-2 block">
              <Label htmlFor="icon" value=" Icon" />
            </div>
            <TextInput
              id="icon"
              value={column.icon}
              onChange={(e) =>
                onChangeColumn({ ...column, icon: e.target.value })
              }
              placeholder=" Icon"
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="name" value="Name" />
            </div>
            <TextInput
              id="name"
              value={column.name}
              onChange={(e) =>
                onChangeColumn({ ...column, name: e.target.value })
              }
              placeholder="Name"
            />
          </div>
          <div>
            <div className="mb-2 block">
              <Label htmlFor="color" value="Color" />
            </div>
            <input
              id="color"
              type="color"
              value={column.color}
              onChange={(e) =>
                onChangeColumn({ ...column, color: e.target.value })
              }
            />
          </div>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button
            onClick={() => {
              updateColumn(projectId!, {
                ...column,
              })
                .then(() => setShowModal(false))
                .catch(toast.error);
            }}
          >
            Save
          </Button>
          <Button color="gray" onClick={() => setShowModal(false)}>
            Cancel
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};
export default ColumnCard;
