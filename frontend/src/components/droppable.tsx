import React from "react";
import { UniqueIdentifier, useDroppable } from "@dnd-kit/core";
import SortableItem from "./sortable";
import { ColumnModel } from "../models/column";
import { TaskModel } from "../models/task";

export interface DroppableProps {
  column: ColumnModel;
  id?: string | null;
  onSelectTask: (task: TaskModel) => void;
  onSelectItem: (id: string) => void;
  onAddItem: () => void;
}

export const Droppable: React.FC<DroppableProps> = ({
  column,
  id,
  onSelectTask,
  onSelectItem,
  onAddItem,
}) => {
  const { isOver, setNodeRef } = useDroppable({
    id: id as UniqueIdentifier,
  });

  const style = {
    backgroundColor: isOver ? "white" : "white",
  };

  return (
    <div
      ref={setNodeRef}
      className="p-4 my-2 rounded-lg shadow w-full min-h-32  "
      style={style}
    >
      {(column.tasks ?? []).length === 0 ? (
        <SortableItem
          onAddItem={onAddItem}
          onSelectItem={onSelectItem}
          isnewitem={true}
          id={`new-item-${column.id}`}
        />
      ) : (
        (column.tasks ?? []).map((item) => (
          <SortableItem
            task={item}
            key={item.id}
            id={item.id!}
            onSelectTask={onSelectTask}
            isnewitem={false}
            onAddItem={onAddItem}
            onSelectItem={onSelectItem}
          />
        ))
      )}
      {(column.tasks ?? []).length > 0 && (
        <div
          className="flex w-full border rounded-lg p-4 text-center flex-row justify-center text-gray-300 mt-4"
          onDoubleClick={onAddItem}
        >
          double click to add item ...
        </div>
      )}
    </div>
  );
};
