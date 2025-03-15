import type { FC } from "react";
import { useEffect, useState } from "react";
import { ProjectModel } from "../models/project";
import { Avatar, Button, Modal, Textarea, TextInput } from "flowbite-react";
import { initial } from "../utils/helper";
import MemberSelectModal from "./MemberSelectModal";
import { getRoles, inviteMember } from "../services/api/commonApi";
import { RoleModel } from "../models/role";
import toast from "react-hot-toast";
import { BsPencil } from "react-icons/bs";
import { updateProject } from "../services/api/projectApi";

interface ProjectHeaderProps {
  project: ProjectModel;
  onChange: (project: ProjectModel) => void;
}

const ProjectHeader: FC<ProjectHeaderProps> = ({ project, onChange }) => {
  const [showModal, setShowModal] = useState(false);
  const [inviteModal, setInviteModal] = useState(false);
  const [inviteEmail, setInviteEmail] = useState("");
  const [inviteFullName, setInviteFullName] = useState("");
  const [inviteRoleId, setInviteRoleId] = useState("");
  const [roles, setRoles] = useState<RoleModel[]>([]);
  const [modalProjectEdit, setModalProjectEdit] = useState(false);

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  useEffect(() => {
    getRoles({ page: 1, size: 10, search: "" }).then((res: any) => {
      setRoles(res.data.items);
    });
  }, []);

  return (
    <div className="h-[80px] flex flex-row justify-between p-4">
      <div className="max-w-[70%] group/item">
        <div className="flex flex-row items-center gap-2">
          <h1 className="text-2xl font-bold">{project?.name}</h1>
          {/* <BsPencil
            className="group/edit invisible group-hover/item:visible text-gray-600 cursor-pointer"
            onClick={() => {
              setModalProjectEdit(true);
            }}
          /> */}
        </div>
        <p className="line-clamp-1">{project?.description}</p>
      </div>
      <div className="flex flex-row gap-4 items-center">
        <Avatar.Group>
          {project?.members?.map((member) => (
            <Avatar
              key={member?.user?.id}
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

        <Button size="xs" onClick={openModal} outline>
          + Member
        </Button>
        <Modal size="4xl" show={showModal} onClose={closeModal}>
          <Modal.Header>Add Member</Modal.Header>
          <Modal.Body>
            <MemberSelectModal
              project={project}
              onInvite={(val) => {
                closeModal();
                setInviteEmail(val);
                setInviteModal(true);
              }}
              // onClose={closeModal}
            />
          </Modal.Body>
          <Modal.Footer className="flex justify-end">
            <Button color="gray" onClick={closeModal}>
              Close
            </Button>
          </Modal.Footer>
        </Modal>
        <Modal
          size="4xl"
          show={inviteModal}
          onClose={() => setInviteModal(false)}
        >
          <Modal.Header>Invite Member</Modal.Header>
          <Modal.Body>
            <form className="flex flex-col gap-4">
              <div>
                <label
                  htmlFor="email"
                  className="block text-sm font-medium text-gray-700"
                >
                  Email
                </label>
                <TextInput
                  id="email"
                  type="email"
                  placeholder="Email"
                  value={inviteEmail}
                  readOnly
                  onChange={(e) => setInviteEmail(e.target.value)}
                />
              </div>
              <div>
                <label
                  htmlFor="full_name"
                  className="block text-sm font-medium text-gray-700"
                >
                  Full Name
                </label>
                <TextInput
                  id="full_name"
                  type="text"
                  placeholder="Full Name"
                  value={inviteFullName}
                  onChange={(e) => {
                    setInviteFullName(e.target.value);
                  }}
                />
              </div>
              <div>
                <label
                  htmlFor="role"
                  className="block text-sm font-medium text-gray-700"
                >
                  Role
                </label>
                <select
                  value={inviteRoleId}
                  id="role"
                  className="block w-full px-4 py-2 text-gray-700 bg-white border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-primary-500 focus:border-primary-500"
                  onChange={(e) => setInviteRoleId(e.target.value)}
                >
                  {roles.map((role) => (
                    <option key={role.id} value={role.id}>
                      {role.name}
                    </option>
                  ))}
                </select>
              </div>
            </form>
          </Modal.Body>
          <Modal.Footer className="flex justify-end">
            <Button
              type="submit"
              color="blue"
              onClick={() => {
                if (!inviteFullName || !inviteRoleId) {
                  alert("Please fill in all fields");
                  return;
                }
                // saveInvite();
                let invitationData = {
                  full_name: inviteFullName,
                  role_id: inviteRoleId,
                  email: inviteEmail,
                  project_id: project?.id,
                };
                inviteMember(invitationData)
                  .then((res: any) => {
                    setInviteModal(false);
                  })
                  .catch(toast.error);
                // setInviteModal(false);
              }}
            >
              Save
            </Button>
            <Button color="gray" onClick={() => setInviteModal(false)}>
              Close
            </Button>
          </Modal.Footer>
        </Modal>
      </div>
      <Modal show={modalProjectEdit} onClose={() => setModalProjectEdit(false)}>
        <Modal.Header>Project Edit</Modal.Header>
        <Modal.Body>
          <form className="flex flex-col gap-4">
            <div>
              <label
                htmlFor="project-name"
                className="block text-sm font-medium text-gray-700"
              >
                Project Name
              </label>
              <TextInput
                id="project-name"
                type="text"
                placeholder="Project Name"
                value={project?.name ?? ""}
                onChange={(e) =>
                  onChange({
                    ...project!,
                    name: e.target.value,
                  })
                }
              />
            </div>
            <div>
              <label
                htmlFor="description"
                className="block text-sm font-medium text-gray-700"
              >
                Description
              </label>
              <Textarea
                rows={10}
                id="description"
                placeholder="Description"
                value={project?.description ?? ""}
                onChange={(e) =>
                  onChange({
                    ...project!,
                    description: e.target.value,
                  })
                }
              />
            </div>
          </form>
        </Modal.Body>
        <Modal.Footer className="flex justify-end">
          <Button
            type="submit"
            color="blue"
            onClick={async () => {
              if (!project?.name || !project?.description) {
                alert("Please fill in all fields");
                return;
              }
              try {
                await updateProject(project!.id!, project!);
                setModalProjectEdit(false);
                toast.success("Project updated successfully");
              } catch (error) {
                toast.error(`${error}`);
              } finally {

              }
            }}
          >
            Save
          </Button>
          <Button color="gray" onClick={() => setModalProjectEdit(false)}>
            Close
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};
export default ProjectHeader;
