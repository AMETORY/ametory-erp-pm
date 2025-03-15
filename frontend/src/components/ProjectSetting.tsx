import type { FC } from "react";
import { ProjectModel } from "../models/project";
import {
  Button,
  Card,
  Datepicker,
  Label,
  Textarea,
  TextInput,
} from "flowbite-react";
import moment from "moment";
import { updateProject } from "../services/api/projectApi";
import toast from "react-hot-toast";

interface ProjectSettingProps {
  project: ProjectModel;
  onChangeProject: (project: ProjectModel) => void;
}

const ProjectSetting: FC<ProjectSettingProps> = ({
  project,
  onChangeProject,
}) => {
  return (
    <div className="grid grid-cols-2 gap-4 px-4">
      <Card>
        <div className="flex space-y-4 flex-col">
          <div className="w-full">
            <Label
              htmlFor="project-name"
              className="block mb-2 text-sm font-bold text-gray-900 dark:text-gray-300"
            >
              Project name
            </Label>
            <TextInput
              type="text"
              id="project-name"
              value={project.name}
              onChange={(e) => {
                onChangeProject({
                  ...project!,
                  name: e.target.value,
                });
              }}
            />
          </div>
          <div className="w-full">
            <Label
              htmlFor="project-name"
              className="block mb-2 text-sm font-bold text-gray-900 dark:text-gray-300"
            >
              Description
            </Label>
            <Textarea
              id="desc"
              rows={5}
              value={project.description}
              onChange={(e) => {
                onChangeProject({
                  ...project!,
                  description: e.target.value,
                });
              }}
            />
          </div>
          <div className="w-full">
            <Label
              htmlFor="project-name"
              className="block mb-2 text-sm font-bold text-gray-900 dark:text-gray-300"
            >
              Deadline
            </Label>
            <Datepicker
              value={moment(project!.deadline).toDate()}
              onChange={(date) => {
                onChangeProject({
                  ...project!,
                  deadline: date!,
                });
              }}
            />
          </div>
          <div>
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
                  toast.success("Project updated successfully");
                } catch (error) {
                  toast.error(`${error}`);
                } finally {
                }
              }}
            >
              Save
            </Button>
          </div>
        </div>
      </Card>
      <Card>
        <h1>Hallo</h1>
      </Card>
   
    </div>
  );
};
export default ProjectSetting;
