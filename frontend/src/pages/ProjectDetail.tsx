import type { FC } from "react";
import React, { useContext, useEffect, useState } from "react";
import {
  DndContext,
  DragOverlay,
  closestCorners,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
  DragEndEvent,
  useDroppable,
  DragStartEvent,
  DragOverEvent,
  UniqueIdentifier,
  closestCenter,
} from "@dnd-kit/core";
import { Droppable } from "../components/droppable";
import {
  SortableContext,
  verticalListSortingStrategy,
  sortableKeyboardCoordinates,
  horizontalListSortingStrategy,
  useSortable,
} from "@dnd-kit/sortable";
import SortableItem from "../components/sortable";
import AdminLayout from "../components/layouts/admin";
import TaskCard from "../components/TaskCard";
import { TaskModel } from "../models/task";
import { ColumnModel } from "../models/column";
import ColumnCard from "../components/ColumnCard";
import { Drawer } from "flowbite-react";
import { BsListTask } from "react-icons/bs";
import { getProject } from "../services/api/projectApi";
import { useParams } from "react-router-dom";
import { ProjectModel } from "../models/project";
import { WebsocketContext } from "../contexts/WebsocketContext";
import { ProfileContext } from "../contexts/ProfileContext";
import ProjectHeader from "../components/ProjectHeader";
interface ProjectDetailProps {}

const ProjectDetail: FC<ProjectDetailProps> = ({}) => {
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const { projectId } = useParams();
  const [activeId, setActiveId] = useState("");
  const [activeColumnId, setActiveColumnId] = useState("");
  const [activeCard, setActiveCard] = useState<TaskModel>();
  const [activeColumn, setActiveColumn] = useState<ColumnModel>();
  const [dragColumn, setDragColumn] = useState(false);
  const [activeTask, setActiveTask] = useState<TaskModel>();
  const [project, setProject] = useState<ProjectModel>();
  const [columns, setColumns] = useState<ColumnModel[]>([]);
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

  useEffect(() => {
    // if (profile?.id != null && profile?.id == wsMsg?.sender_id) {
      // console.log("wsMsg", wsMsg);
    // }
  }, [wsMsg, profile]);

  useEffect(() => {
    if (!projectId) return;
    getProject(projectId).then((resp: any) => {
      setProject(resp.data);
      setColumns(
        resp.data.columns.map((column: any) => {
          return {
            ...column,
            tasks: column.tasks ?? [],
          };
        })
      );
    });
  }, []);

  const handleDragEnd = (event: DragEndEvent) => {
    setActiveId("");
    setActiveColumnId("");
    setActiveCard(undefined);
    const { active, over } = event;
    const { id } = active;
    let isColumn: boolean =
      active.data.current?.sortable?.containerId?.includes("Sortable");
    setDragColumn(isColumn);
    // console.log(event);
    if (event.over) {
      if (!isColumn) {
        const activeColumn = columns.find(
          (column) => column.id === active.data.current?.sortable?.containerId
        );

        const overColumn = columns.find(
          (column) => column.id === over?.data.current?.sortable?.containerId
        );

        // console.log(activeColumn, overColumn);
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
        }
      } else {
        const activeIndex = columns.findIndex((item) => item.id == id);
        const overIndex = columns.findIndex((item) => item.id == over?.id);
        console.log(activeIndex, overIndex);
        if (
          activeIndex !== -1 &&
          overIndex !== -1 &&
          activeIndex !== overIndex
        ) {
          const movedColumn = columns.splice(activeIndex, 1)[0];
          columns.splice(overIndex, 0, movedColumn);
          setColumns([...columns]);
        }
      }
    }
  };

  const handleDragOver = (event: DragOverEvent) => {
    const { active, over } = event;
    const { id } = active;
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

  return (
    <AdminLayout>
      {project && <ProjectHeader project={project} />}
      <div className="p-4 h-full overflow-x-scroll unselected">
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
            <div className=" flex-nowrap flex gap-4 w-fit">
              {columns.map((column) => {
                return (
                  <ColumnCard
                    projectId={projectId!}
                    key={column.id}
                    column={column}
                    columns={columns}
                    onChange={setColumns}
                    onSelectTask={(val) => {
                      setActiveTask(val);
                      // console.log(val);
                    }}
                  />
                );
              })}
              <div
                className="flex flex-col items-center "
                style={{ width: 300 }}
              >
                <div
                  className="border p-4 rounded-lg text-center cursor-pointer hover:bg-gray-50 transform w-full"
                  onClick={() => {
                    setColumns([
                      ...columns,
                      {
                        id: `column${columns.length + 1}`,
                        name: "New Column",
                        order: columns.length,
                        color: `hsl(${Math.floor(
                          Math.random() * 360
                        )}, 100%, 90%)`,
                        tasks: [],
                      },
                    ]);
                  }}
                >
                  Add New Column
                </div>
              </div>
            </div>
            <DragOverlay>
              {dragColumn && activeColumn ? (
                <ColumnCard
                  projectId={projectId!}
                  column={activeColumn}
                  columns={columns}
                  onChange={setColumns}
                  onSelectTask={(val) => {}}
                />
              ) : (
                <TaskCard
                  onClick={(val) => {
                    console.log(val);
                  }}
                  task={activeCard as TaskModel}
                  isdragged={true}
                />
              )}
            </DragOverlay>
          </SortableContext>
        </DndContext>
      </div>
      <Drawer
        style={{ width: 1000 }}
        open={activeTask !== undefined}
        onClose={() => setActiveTask(undefined)}
        position="right"
      >
        <Drawer.Header titleIcon={BsListTask} title={activeTask?.name} />
        <Drawer.Items className="pt-4">{activeTask?.name}</Drawer.Items>
      </Drawer>
    </AdminLayout>
  );
};
export default ProjectDetail;
