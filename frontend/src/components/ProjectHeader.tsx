import type { FC } from "react";
import { useState } from "react";
import { ProjectModel } from "../models/project";
import { Avatar, Button, Modal } from "flowbite-react";
import { initial } from "../utils/helper";
import MemberSelectModal from "./MemberSelectModal";

interface ProjectHeaderProps {
  project: ProjectModel;
}

const ProjectHeader: FC<ProjectHeaderProps> = ({ project }) => {
  const [showModal, setShowModal] = useState(false);

  const openModal = () => setShowModal(true);
  const closeModal = () => setShowModal(false);

  return (
    <div className="h-[80px] flex flex-row justify-between p-4">
      <div>
        <h1 className="text-2xl font-bold">{project?.name}</h1>
        <p>{project?.description}</p>
      </div>
      <div className="flex flex-row gap-2 items-center">
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
        <Button onClick={openModal} outline>+ Member</Button>
        <Modal
          show={showModal}
          onClose={closeModal}
        >
          <Modal.Header>
            Add Member
          </Modal.Header>
          <Modal.Body>
            <MemberSelectModal
              project={project}
              // onClose={closeModal}
            />
          </Modal.Body>
        </Modal>
      </div>
    </div>
  );
};
export default ProjectHeader;

