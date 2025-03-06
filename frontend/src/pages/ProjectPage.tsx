import type { FC } from "react";
import AdminLayout from "../components/layouts/admin";
import {
  Avatar,
  Badge,
  Button,
  Datepicker,
  Drawer,
  Modal,
  Table,
} from "flowbite-react";
import { useContext, useEffect, useState } from "react";
import { Industry, IndustryColumn, ProjectModel } from "../models/project";
import { useNavigate } from "react-router-dom";
import {
  createProject,
  deleteProject,
  getProjects,
  getProjectTemplates,
} from "../services/api/projectApi";
import { LoadingContext } from "../contexts/LoadingContext";
import Moment from "react-moment";
import { PaginationResponse } from "../objects/pagination";
import { getPagination, initial } from "../utils/helper";
import toast from "react-hot-toast";

interface ProjectPageProps {}

const ProjectPage: FC<ProjectPageProps> = ({}) => {
  const [showModal, setShowModal] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const [projects, setProjects] = useState<ProjectModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(10);
  const [search, setSearch] = useState("");
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [project, setProject] = useState<ProjectModel | null>(null);
  const [industries, setIndustries] = useState<Industry[]>([]);
  const [selectedIndustry, setSelectedIndustry] = useState<Industry>();

  const navigate = useNavigate();
  const handleCreateProject = async () => {
    try {
      if (!project) return;
      setLoading(true);

      const newProject = await createProject({
        name: project.name,
        description: project.description,
        deadline: project.deadline,
        status: project.status,
        columns: selectedIndustry?.columns,
      });
      getAllProjects();
      setShowModal(false);
    } catch (error: any) {
      toast.error("Error creating project:", error);
      alert(error.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    getAllProjects();
    getProjectTemplates()
      .then((response: any) => {
        setIndustries(response.data);
      })
      .catch((error) => {
        toast.error("Error fetching industries:", error);
      });
    return () => {};
  }, []);

  const getAllProjects = async () => {
    try {
      setLoading(true);
      const resp: any = await getProjects({ page, size, search });
      setProjects(resp.data.items);
      setPagination(getPagination(resp.data));
    } catch (error) {
      toast.error(`${error}`);
    } finally {
      setLoading(false);
    }
  };
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Project</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setProject({
                name: "",
                deadline: new Date(),
                tasks: [],
                columns: [],
              });
              setShowModal(true);
            }}
          >
            + Create new project
          </Button>
        </div>
        <Table>
          <Table.Head>
            <Table.HeadCell>Project Name</Table.HeadCell>
            <Table.HeadCell>Deadline</Table.HeadCell>
            <Table.HeadCell>Status</Table.HeadCell>
            <Table.HeadCell>Assigned Member</Table.HeadCell>
            <Table.HeadCell></Table.HeadCell>
          </Table.Head>

          <Table.Body className="divide-y">
            {projects.length === 0 && (
              <Table.Row>
                <Table.Cell colSpan={5} className="text-center">
                  No projects found.
                </Table.Cell>
              </Table.Row>
            )}
            {projects.map((project) => (
              <Table.Row
                key={project.id}
                className="bg-white dark:border-gray-700 dark:bg-gray-800"
              >
                <Table.Cell
                  className="whitespace-nowrap font-medium text-gray-900 dark:text-white cursor-pointer hover:font-semibold"
                  onClick={() => navigate(`/project/${project.id}`)}
                >
                  {project.name}
                </Table.Cell>
                <Table.Cell>
                  {project.deadline && (
                    <Moment format="YYYY-MM-DD" date={project.deadline} />
                  )}
                </Table.Cell>
                <Table.Cell>
                  <div className="w-fit">
                    <Badge>Active</Badge>
                  </div>
                </Table.Cell>
                <Table.Cell>
                  <Avatar.Group>
                    {project?.members?.map((member) => (
                      <Avatar
                        key={member.id}
                        size="xs"
                        img={member?.user?.profile_picture?.url}
                        rounded
                        stacked
                        placeholderInitials={initial(member?.user?.full_name)}
                      />
                    ))}
                    {(project?.members ?? []).length > 5 && (
                      <Avatar.Counter
                        total={(project?.members ?? []).length - 5}
                        href="#"
                      />
                    )}
                  </Avatar.Group>
                </Table.Cell>
                <Table.Cell>
                  {/* <a
                    href="#"
                    className="font-medium text-cyan-600 hover:underline dark:text-cyan-500"
                  >
                    Edit
                  </a> */}
                  <a
                    href="#"
                    className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2"
                    onClick={(e) => {
                      e.preventDefault();
                      if (
                        window.confirm(
                          `Are you sure you want to delete project ${project.name}?`
                        )
                      ) {
                        deleteProject(project?.id!).then(() => {
                          getAllProjects();
                        });
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

        <Modal
          show={showModal}
          onClose={() => {
            setShowModal(false);
          }}
        >
          <Modal.Header>Create new project</Modal.Header>
          <Modal.Body>
            <form
              onSubmit={(e) => {
                e.preventDefault();
                // const formData = new FormData(e.currentTarget);
                // const project: ProjectModel = {
                //   name: formData.get("name") as string,
                //   deadline: formData.get("deadline") as string,
                //   status: formData.get("status") as string,
                //   assignedMember: formData.get("assignedMember") as string,
                // };
                // handleCreateProject(project);
              }}
              className="space-y-6"
            >
              <div className="mb-4">
                <label
                  htmlFor="name"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                >
                  Name
                </label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="Project name"
                  value={project?.name}
                  onChange={(e) =>
                    setProject((prev) => ({ ...prev, name: e.target.value }))
                  }
                  required
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="description"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                >
                  Description
                </label>
                <textarea
                  id="description"
                  name="description"
                  className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  placeholder="Project description"
                  value={project?.description}
                  onChange={(e) =>
                    setProject((prev) => ({
                      ...prev,
                      description: e.target.value,
                    }))
                  }
                  rows={3}
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="deadline"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                >
                  Deadline
                </label>

                <Datepicker
                  placeholder="Deadline"
                  value={project?.deadline}
                  onChange={(date) =>
                    setProject((prev) => ({ ...prev, deadline: date }))
                  }
                  required
                  id="deadline"
                />
              </div>
              <div className="mb-4">
                <label
                  htmlFor="status"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                >
                  Status
                </label>
                <select
                  id="status"
                  name="status"
                  className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  value={project?.status}
                  onChange={(e) =>
                    setProject((prev) => ({ ...prev, status: e.target.value }))
                  }
                  required
                >
                  <option value="Active">Active</option>
                  <option value="Completed">Completed</option>
                  <option value="Pending">Pending</option>
                </select>
              </div>

              <div className="mb-4">
                <label
                  htmlFor="assignedMember"
                  className="block mb-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                >
                  Template
                </label>
                <select
                  id="template"
                  name="template"
                  className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                  onChange={(e) => {
                    setSelectedIndustry(
                      industries.find((i) => i.industry === e.target.value)
                    );
                  }}
                  required
                >
                  <option value={""}>Select Template</option>
                  {industries.map((val) => (
                    <option key={val.industry} value={val.industry}>
                      {val.industry}
                    </option>
                  ))}
                </select>
              </div>
              <div className="flex justify-end">
                <button
                  type="submit"
                  className="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
                  onClick={() => handleCreateProject()}
                >
                  Create project
                </button>
              </div>
            </form>
          </Modal.Body>
        </Modal>
      </div>
      
    </AdminLayout>
  );
};
export default ProjectPage;
