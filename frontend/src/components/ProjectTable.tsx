import { useState, type FC } from "react";
import { ColumnModel } from "../models/column";
import { ProjectModel } from "../models/project";
import {
  DndContext,
  DragEndEvent,
  DragOverEvent,
  DragOverlay,
  DragStartEvent,
  KeyboardSensor,
  PointerSensor,
  UniqueIdentifier,
  useSensor,
  useSensors,
} from "@dnd-kit/core";
import {
  horizontalListSortingStrategy,
  SortableContext,
  sortableKeyboardCoordinates,
} from "@dnd-kit/sortable";
import { TaskModel } from "../models/task";
import Moment from "react-moment";
import { Avatar } from "flowbite-react";
import { initial } from "../utils/helper";
import ColumnTable from "./ColumnTable";
import { getTask, moveTask, rearrangeTask } from "../services/api/taskApi";
import toast from "react-hot-toast";
import TaskTable from "./TaskTable";
import { rearrangeColumns } from "../services/api/projectApi";

interface ProjectTableProps {
  project: ProjectModel;
  columns: ColumnModel[];
  setColumns: (columns: ColumnModel[]) => void;
  setActiveTask: (task: TaskModel) => void;
}

const ProjectTable: FC<ProjectTableProps> = ({
  project,
  columns,
  setColumns,
  setActiveTask,
}) => {
  const [activeId, setActiveId] = useState("");
  const [activeColumnId, setActiveColumnId] = useState("");
  const [activeCard, setActiveCard] = useState<TaskModel>();
  const [activeColumn, setActiveColumn] = useState<ColumnModel>();
  const [dragColumn, setDragColumn] = useState(false);

  const handleDragEnd = (event: DragEndEvent) => {
    setActiveId("");
    setActiveColumnId("");
    setActiveCard(undefined);
    const { active, over } = event;
    const { id } = active;
    let isColumn: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");
    setDragColumn(isColumn);
    if (event.over) {
      if (!isColumn) {
        const activeColumn = columns.find(
          (column) => column.id === active.data.current?.sortable?.containerId
        );

        const overColumn = columns.find(
          (column) => column.id === over?.data.current?.sortable?.containerId
        );

        if (activeColumn && overColumn) {
          const activeIndex = (activeColumn.tasks ?? []).findIndex(
            (item) => item.id === id
          );
          const overIndex = (overColumn.tasks ?? []).findIndex(
            (item) => item.id === over?.id
          );
          const item = (activeColumn.tasks ?? []).splice(activeIndex, 1)[0];
          (overColumn.tasks ?? []).splice(overIndex, 0, item);
          // console.log(item);
          // Reload the columns to trigger a re-render

          setColumns([
            ...columns.slice(0, columns.indexOf(activeColumn)),
            activeColumn,
            ...columns.slice(columns.indexOf(activeColumn) + 1),
          ]);
          if (overColumn.id != activeColumn.id) {
            moveTask(project!.id!, item!.id!, {
              column_id: overColumn.id,
              source_column_id: activeColumn.id,
              order_number:
                overColumn.tasks?.findIndex((over) => over.id == item?.id) ??
                0 + 1,
            }).catch(toast.error);
          } else {
            rearrangeTask(project!.id!, activeColumn).catch(toast.error);
          }
        }
      } else {
        const activeIndex = columns.findIndex((item) => item.id == id);
        const overIndex = columns.findIndex((item) => item.id == over?.id);
        // console.log(activeIndex, overIndex);
        if (
          activeIndex !== -1 &&
          overIndex !== -1 &&
          activeIndex !== overIndex
        ) {
          let columnsBefore = [...columns];
          const movedColumn = columns.splice(activeIndex, 1)[0];
          columns.splice(overIndex, 0, movedColumn);
          setColumns([...columns]);

          console.log(movedColumn);
          rearrangeColumns(project!.id!, {
            columns: [...columns],
          }).catch((err) => {
            // console.error(err);
            toast.error(err.message || err.error);
            setColumns(columnsBefore);
          });
        }
      }
    }
  };

  const handleDragStart = (event: DragStartEvent) => {
    const { active } = event;
    const { id } = active;
    let isColumn: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");
    setDragColumn(isColumn);

    if (!isColumn) {
      let columnId = event.active.data.current?.sortable?.containerId as string;
      let itemId = id as string;

      setActiveColumnId(
        event.active.data.current?.sortable?.containerId as string
      );
      setActiveId(id as string);
      setActiveCard(
        (columns.find((column) => column.id === columnId)?.tasks ?? []).find(
          (item) => item.id === itemId
        )
      );
    } else {
      const activeIndex = columns.findIndex((item) => item.id == id);
      setActiveColumn(columns[activeIndex]);
    }
  };
  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    const { id } = active;
  };
  const sensors = useSensors(
    useSensor(PointerSensor, {
      activationConstraint: {
        delay: 200,
        tolerance: 0,
      },
    }),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  return (
    <DndContext
      onDragEnd={handleDragEnd}
      onDragStart={handleDragStart}
      onDragOver={handleDragOver}
      sensors={sensors}
    >
      <SortableContext
        items={columns.map((column) => ({
          id: column.id as UniqueIdentifier,
        }))}
        strategy={horizontalListSortingStrategy}
      >
        <div className="w-full px-4 space-y-4">
          {columns.map((column, i) => (
            <ColumnTable
              key={i}
              projectId={project!.id!}
              column={column}
              columns={columns}
              onChange={setColumns}
              onChangeColumn={(val) => {
                setColumns([
                  ...columns.map((c) => {
                    if (c.id === val.id) {
                      return val;
                    }
                    return c;
                  }),
                ]);
              }}
              onSelectTask={(val) => {
                setActiveTask(val);
                // console.log(val);
              }}
              onAddItem={(val) => {
                getTask(project!.id!, val)
                  .then((resp: any) => setActiveTask(resp.data))
                  .catch(toast.error);
              }}
            />
          ))}
        </div>
        <DragOverlay>
          {dragColumn && activeColumn ? (
            <ColumnTable
              projectId={project!.id!}
              column={activeColumn!}
              columns={columns}
              onChange={setColumns}
              onChangeColumn={(val) => {}}
              onSelectTask={(val) => {
                // console.log(val);
              }}
              onAddItem={(val) => {}}
            />
          ) : (
            <TaskTable task={activeCard!} onSelectTask={(val) => {}} />
          )}
        </DragOverlay>
      </SortableContext>
    </DndContext>
  );
};
export default ProjectTable;
