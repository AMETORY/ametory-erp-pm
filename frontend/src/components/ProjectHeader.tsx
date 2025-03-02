import type { FC } from "react";
import { ProjectModel } from "../models/project";
import { Avatar, Button } from "flowbite-react";
import { initial } from "../utils/helper";

interface ProjectHeaderProps {
  project: ProjectModel;
}

const ProjectHeader: FC<ProjectHeaderProps> = ({ project }) => {
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
        <Button outline>+ Member</Button>
      </div>
    </div>
  );
};
export default ProjectHeader;
