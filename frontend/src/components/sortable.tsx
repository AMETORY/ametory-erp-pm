import React from "react";
import { useSortable } from "@dnd-kit/sortable";
import { CSS } from "@dnd-kit/utilities";
import TaskCard from "./TaskCard";
import type { FC } from "react";

interface SortableItemProps {
  id: string;
  isnewitem: boolean;
  task?: any; // Replace 'any' with the actual type of task if available
  onAddItem: () => void;
  onSelectItem: (id: string) => void; // Replace 'any' with the actual type of id if available
  onSelectTask?: (task: any) => void; // Replace 'any' with the actual type of task if available
}

const SortableItem: FC<SortableItemProps> = ({
  id,
  isnewitem,
  task,
  onAddItem,
  onSelectTask,
  onSelectItem,
  ...props
}) => {
  const { attributes, listeners, setNodeRef, transform, transition } = useSortable({
    id: id,
    disabled: isnewitem,
  });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
  };

  return (
    <div ref={setNodeRef} style={style} {...attributes} {...listeners} {...props}>
      {isnewitem ? (
        <div
          className="flex w-full border rounded-lg p-4 text-center flex-row justify-center text-gray-300"
          onDoubleClick={onAddItem}
        >
          Drop Here or double click to add item ...
        </div>
      ) : (
        <TaskCard task={task} onClick={() => onSelectTask?.(task)} isdragged={false} />
      )}
    </div>
  );
};

export default SortableItem;

