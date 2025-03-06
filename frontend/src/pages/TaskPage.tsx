import { useContext, useEffect, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Avatar,
  Badge,
  Button,
  Drawer,
  Pagination,
  Table,
  Tabs,
} from "flowbite-react";
import { PaginationResponse } from "../objects/pagination";
import { LoadingContext } from "../contexts/LoadingContext";
import { TaskModel } from "../models/task";
import toast, { Toaster } from "react-hot-toast";
import {
  deleteTask,
  getMyTasks,
  getMyWatchedTasks,
} from "../services/api/taskApi";
import { getPagination, initial } from "../utils/helper";
import Moment from "react-moment";
import { useNavigate } from "react-router-dom";
import TaskDetail from "../components/TaskDetail";
import { BsListTask } from "react-icons/bs";
import { ProjectModel } from "../models/project";
import { getProject } from "../services/api/projectApi";

interface TaskPageProps {}

const TaskPage: FC<TaskPageProps> = ({}) => {
  const navigate = useNavigate();
  const { loading, setLoading } = useContext(LoadingContext);
  const [mounted, setMounted] = useState(false);
  const [page, setPage] = useState(1);
  const [size, setSize] = useState(10);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [pageWatched, setPageWatched] = useState(1);
  const [sizeWatched, setSizeWatched] = useState(10);
  const [searchWatched, setSearchWatched] = useState("");
  const [paginationWatched, setPaginationWatched] =
    useState<PaginationResponse>();
  const [myTask, setMyTask] = useState<TaskModel[]>([]);
  const [watchedTask, setWatchedTask] = useState<TaskModel[]>([]);
  const [activeTask, setActiveTask] = useState<TaskModel>();
  const [isTaskFullScreen, setIsTaskFullScreen] = useState(false);
  const [project, setProject] = useState<ProjectModel>();

  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted) {
      getAllMyTasks();
    }
  }, [mounted, page, size, search]);
  useEffect(() => {
    if (mounted) {
      getAllWatchedTasks();
    }
  }, [mounted, pageWatched, sizeWatched, searchWatched]);

  const getAllMyTasks = async () => {
    try {
      setLoading(true);
      let resp: any = await getMyTasks({ page, size, search });
      setMyTask(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  const getAllWatchedTasks = async () => {
    try {
      setLoading(true);
      let resp: any = await getMyWatchedTasks({
        page: pageWatched,
        size: sizeWatched,
        search: searchWatched,
      });
      setWatchedTask(resp.data.items);
      setPaginationWatched(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };

  const getDetail = (t: TaskModel) => {
    setActiveTask(t);
    setLoading(true);
    getProject(t.project_id!)
      .then((v: any) => setProject(v.data))
      .catch(toast.error)
      .finally(() => setLoading(false));
  };
  const renderTable = (tasks: TaskModel[], mode: string = "my") => (
    <div>
      <Table>
        <Table.Head>
          <Table.HeadCell>Task</Table.HeadCell>
          <Table.HeadCell>Assigned To</Table.HeadCell>
          <Table.HeadCell>Watcher</Table.HeadCell>
          <Table.HeadCell>Status</Table.HeadCell>
          <Table.HeadCell></Table.HeadCell>
        </Table.Head>
        <Table.Body className="divide-y">
          {tasks.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={5} className="text-center">
                No projects found.
              </Table.Cell>
            </Table.Row>
          )}
          {tasks.map((task) => (
            <Table.Row
              key={task.id}
              className="bg-white dark:border-gray-700 dark:bg-gray-800"
            >
              <Table.Cell
                className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                onClick={() => {
                  getDetail(task);
                }}
              >
                {task.name}
              </Table.Cell>
              <Table.Cell>
                {task?.assignee && (
                  <div className="flex flex-row gap-2">
                    <Avatar
                      key={task?.assignee.id}
                      size="xs"
                      img={task?.assignee?.user?.profile_picture?.url}
                      rounded
                      stacked
                      placeholderInitials={initial(
                        task?.assignee?.user?.full_name
                      )}
                    />
                    {task?.assignee?.user?.full_name}
                  </div>
                )}
              </Table.Cell>
              {/* <Table.Cell>
              {task.end_date && (
                <Moment format="YYYY-MM-DD" date={task.end_date} />
              )}
            </Table.Cell> */}

              <Table.Cell>
                <Avatar.Group>
                  {task?.watchers?.map((member) => (
                    <Avatar
                      key={member.id}
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
              </Table.Cell>
              <Table.Cell>
                <div className="w-fit">
                  {task?.status == "ACTIVE" && (
                    <Badge color="blue">{task?.status}</Badge>
                  )}
                  {task?.status == "COMPLETED" && (
                    <Badge color="green">{task?.status}</Badge>
                  )}
                </div>
              </Table.Cell>
              <Table.Cell>
                <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                    onClick={() => getDetail(task)}
                  >
                    View
                  </a>
                <a
                  href="#"
                  className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                  onClick={(e) => {
                    e.preventDefault();
                    if (
                      window.confirm(
                        `Are you sure you want to delete taks ${task.name}?`
                      )
                    ) {
                      deleteTask(task?.id!).then(() => {});
                    }
                  }}
                >
                  Delete
                </a>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
      <Pagination
      className="mt-4"
        currentPage={mode == "my" ? page : pageWatched}
        totalPages={
          mode == "my"
            ? (pagination?.total_pages ?? 0)
            : (paginationWatched?.total_pages ?? 0)
        }
        onPageChange={(val) => {
          if (mode == "my") {
            setPage(val);
          } else {
            setPageWatched(val);
          }
        }}
        showIcons
      />
    </div>
  );
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Task</h1>
        </div>
        <Tabs aria-label="Pills" variant="pills">
          <Tabs.Item active title="My Task">
            {renderTable(myTask)}
          </Tabs.Item>
          <Tabs.Item title="Watched">{renderTable(watchedTask)}</Tabs.Item>
        </Tabs>
      </div>
      {activeTask && project && (
        <Drawer
          style={{ width: !isTaskFullScreen ? "1000px" : "100%" }}
          open={activeTask !== undefined}
          onClose={() => setActiveTask(undefined)}
          position="right"
        >
          <Drawer.Header titleIcon={BsListTask} title={activeTask?.name} />
          <Drawer.Items
            className="pt-4  "
            style={{ height: "calc(100vh - 70px)" }}
          >
            <TaskDetail
              task={activeTask}
              project={project}
              onSwitchFullscreen={() => setIsTaskFullScreen((val) => !val)}
            />
          </Drawer.Items>
        </Drawer>
      )}
    </AdminLayout>
  );
};
export default TaskPage;
