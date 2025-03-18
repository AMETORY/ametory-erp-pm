import type { FC } from "react";
import { ProjectModel, ProjectPreference } from "../models/project";
import {
  Button,
  Card,
  Datepicker,
  Label,
  Textarea,
  TextInput,
  ToggleSwitch,
} from "flowbite-react";
import moment from "moment";
import { updateProject, updateProjectPreference } from "../services/api/projectApi";
import toast from "react-hot-toast";

interface ProjectSettingProps {
  project: ProjectModel;
  preference: ProjectPreference;
  onChangeProject: (project: ProjectModel) => void;
  onChangePreference: (preference: ProjectPreference) => void;
}

const ProjectSetting: FC<ProjectSettingProps> = ({
  project,
  preference,
  onChangeProject,
  onChangePreference,
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
        <div className="flex flex-col h-full space-y-8">
          <h3 className=" text-lg font-semibold">Preferences</h3>
          <div className="flex justify-between item">
            <Label className="text-md">Rapid API</Label>
            <ToggleSwitch
              checked={preference!.rapid_api_enabled}
              onChange={(enabled) => {
                onChangePreference({
                  ...preference!,
                  rapid_api_enabled: enabled,
                });
              }}
            />
          </div>
          <div className="flex justify-between item">
            <Label className="text-md">Contact</Label>
            <ToggleSwitch
              checked={preference!.contact_enabled}
              onChange={(enabled) => {
                onChangePreference({
                  ...preference!,
                  contact_enabled: enabled,
                });
              }}
            />
          </div>
          <div className="flex justify-between item">
            <Label className="text-md">Custom Attribute</Label>
            <ToggleSwitch
              checked={preference!.custom_attribute_enabled}
              onChange={(enabled) => {
                onChangePreference({
                  ...preference!,
                  custom_attribute_enabled: enabled,
                });
              }}
            />
          </div>
          <div className="flex justify-between item">
            <Label className="text-md">Gemini</Label>
            <ToggleSwitch
              checked={preference!.gemini_enabled}
              onChange={(enabled) => {
                onChangePreference({
                  ...preference!,
                  gemini_enabled: enabled,
                });
              }}
            />
          </div>
          <div className="flex justify-between item">
            <Label className="text-md">Form</Label>
            <ToggleSwitch
              checked={preference!.form_enabled}
              onChange={(enabled) => {
                onChangePreference({
                  ...preference!,
                  form_enabled: enabled,
                });
              }}
            />
          </div>
        </div>
        <div>
          <Button
            type="submit"
            color="blue"
            onClick={async () => {
        
              try {
                await updateProjectPreference(project!.id!, preference!);
                toast.success("Project Preference updated successfully");
              } catch (error) {
                toast.error(`${error}`);
              } finally {
              }
            }}
          >
            Save
          </Button>
        </div>
      </Card>
    </div>
  );
};
export default ProjectSetting;
