import { useContext, useEffect, useState, type FC } from "react";
import { ProjectActivityModel, ProjectModel } from "../models/project";
import { Card } from "flowbite-react";
import Chart from "react-google-charts";
import {
  countCompletedTasks,
  countCreatedTasks,
  countDueTasks,
  countTasksByColumn,
  countTasksByPriority,
  countTasksBySeverity,
  countUpdatedTasks,
  getRecentActivities,
} from "../services/api/projectApi";
import Moment from "react-moment";
import { WebsocketContext } from "../contexts/WebsocketContext";

interface ProjectSummaryProps {
  project: ProjectModel;
}

const ProjectSummary: FC<ProjectSummaryProps> = ({ project }) => {
  const [mounted, setMounted] = useState(false);
  const [done, setDone] = useState(0);
  const [created, setCreated] = useState(0);
  const [due, setDue] = useState(0);
  const [updated, setUpdated] = useState(0);
  const [rangeDays, setRangeDays] = useState(7);
  const [activities, setActivities] = useState<ProjectActivityModel[]>([]);
  const [columnChartData, setColumnChartData] = useState<any[][]>([]);
  const [priorityChartData, setPriorityChartData] = useState<any[][]>([]);
  const [severityChartData, setSeverityChartData] = useState<any[][]>([]);
  const { isWsConnected, setWsConnected, wsMsg, setWsMsg } =
    useContext(WebsocketContext);
  useEffect(() => {
    setMounted(true);
  }, []);

  useEffect(() => {
    if (mounted && project) {
      getAllCharts();
    }
  }, [mounted, project]);

  useEffect(() => {
    if (wsMsg?.project_id == project?.id) {
      getAllCharts();
    }
  }, [wsMsg, project]);

  const getAllCharts = () => {
    countCompletedTasks(project.id!, 7).then((e: any) => setDone(e.count));
    countUpdatedTasks(project.id!, 7).then((e: any) => setUpdated(e.count));
    countCreatedTasks(project.id!, 7).then((e: any) => setCreated(e.count));
    countDueTasks(project.id!, 7).then((e: any) => setDue(e.count));
    getRecentActivities(project.id!).then((e: any) => setActivities(e.data));
    countTasksByPriority(project.id!).then((e: any) => {
      let chartData: any[][] = [["Priority", "Count"]];
      for (const element of e.data) {
        chartData.push([element.label, element.count]);
      }

      setPriorityChartData(chartData);
    });
    countTasksBySeverity(project.id!).then((e: any) => {
      let chartData: any[][] = [["Severity", "Count"]];
      for (const element of e.data) {
        chartData.push([element.label, element.count]);
      }

      setSeverityChartData(chartData);
    });
    countTasksByColumn(project.id!).then((e: any) => {
      let chartData: any[][] = [["Column", "Count"]];
      for (const element of e.data) {
        chartData.push([element.name, element.count_tasks]);
      }

      setColumnChartData(chartData);
    });
  };

  return (
    <div
      className=" px-4 flex flex-col overflow-y-auto"
      style={{ height: "calc(100vh - 240px)" }}
    >
      <div className="w-full mx-auto space-y-8">
        <div>
          <h1 className="text-2xl font-bold  text-center mt-8">
            Good morning, Rahmat Supriatna
          </h1>
          <p className=" text-center">
            Here’s a summary of your project’s status, priorities, workload, and
            more.
          </p>
        </div>
        <div className="grid grid-cols-2 sm:grid-cols-4 gap-4">
          <Card>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {done} Done
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last {rangeDays} days
            </p>
          </Card>
          <Card>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {updated} Updated
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last {rangeDays} days
            </p>
          </Card>
          <Card>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {created} Created
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last {rangeDays} days
            </p>
          </Card>
          <Card>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              {due} Due
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last {rangeDays} days
            </p>
          </Card>
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <Card className="h-[420px]">
            <Chart
              style={{ borderRadius: "8px" }}
              chartType="PieChart"
              data={columnChartData}
              options={{
                is3D: true,
                title: "Tasks By Stage",
              }}
              height={"100%"}
            />
          </Card>
          <Card className="h-[420px]">
            <div className="flex flex-col h-full overflow-y-auto">
              <h3 className="font-bold text-lg">Recent Activities</h3>
              <ul className="space-y-4">
                {activities?.map((activity) => (
                  <li key={activity.id} className="">
                    <span className="text-gray-600 dark:text-gray-300 hover:font-semibold">
                      {activity.member?.user?.full_name}
                    </span>{" "}
                    <strong>
                      {activity.activity_type?.replaceAll("_", " ")}
                    </strong>{" "}
                    at <Moment fromNow>{activity.activity_date}</Moment>
                  </li>
                ))}
              </ul>
            </div>
          </Card>
          <Card className="h-[420px]">
            <Chart
              style={{ borderRadius: "8px" }}
              chartType="PieChart"
              data={priorityChartData}
              options={{
                title: "Priority Tasks",
                colors: ["#8BC34A", "#F7DC6F", "#FFC107", "#F44336"],
              }}
            />
            <Chart
              style={{ borderRadius: "8px" }}
              chartType="PieChart"
              data={severityChartData}
              options={{
                title: "Severity Tasks",
                colors: ["#8BC34A", "#F7DC6F", "#FFC107", "#F44336"],
              }}
            />
          </Card>
        </div>
      </div>
    </div>
  );
};
export default ProjectSummary;
