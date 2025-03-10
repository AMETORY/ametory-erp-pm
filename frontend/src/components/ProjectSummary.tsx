import type { FC } from "react";
import { ProjectModel } from "../models/project";
import { Card } from "flowbite-react";
import Chart from "react-google-charts";

interface ProjectSummaryProps {
  project: ProjectModel;
}

const ProjectSummary: FC<ProjectSummaryProps> = ({ project }) => {
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
          <Card href="#">
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              0 Done
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last 7 days
            </p>
          </Card>
          <Card href="#">
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              0 Updated
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last 7 days
            </p>
          </Card>
          <Card href="#">
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              0 Created
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last 7 days
            </p>
          </Card>
          <Card href="#">
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
              0 Due
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
              in last 7 days
            </p>
          </Card>
        </div>
        <div className="grid grid-cols-1 sm:grid-cols-3 gap-4">
          <Card className="h-[420px]" href="#">
            <Chart
              style={{ borderRadius: "8px" }}
              chartType="PieChart"
              data={[
                ["Task", "Hours per Day"],
                ["Work", 9],
                ["Eat", 2],
                ["Commute", 2],
                ["Watch TV", 2],
                ["Sleep", 7],
              ]}
              options={{
                is3D: true,
                title: "My Daily Activities",
              }}
              height={"100%"}
            />
          </Card>
          <Card className="h-[420px]" href="#">
            <div className="flex flex-col h-full">
              <h3 className="font-bold text-lg">Recent Activities</h3>
              <ul className="mt-4">
                <li>
                  <strong>ADD COMMENT</strong> Rahmat Supriatna{" "}
                  <em>at 3 days ago</em>
                </li>
                <li>
                  <strong>CREATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 3 days ago</em>
                </li>
                <li>
                  <strong>ADD COMMENT</strong> Tatang Kalan{" "}
                  <em>at 11 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 11 hours ago</em>
                </li>
                <li>
                  <strong>ADD COMMENT</strong> Tatang Kalan{" "}
                  <em>at 11 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
                <li>
                  <strong>UPDATE TASK</strong> Rahmat Supriatna{" "}
                  <em>at 8 hours ago</em>
                </li>
              </ul>
            </div>
          </Card>
          <Card className="h-[420px]" href="#">
            <Chart
              chartType="ColumnChart"
              width="100%"
              height="100%"
              data={[
                ["Year", "Sales", "Expenses"],
                ["2013", 1000, 400],
                ["2014", 1170, 460],
                ["2015", 660, 1120],
                ["2016", 1030, 540],
              ]}
              options={{
                title: "Company Performance",

                bar: { groupWidth: "75%" },
                legend: { position: "bottom" },
              }}
            />
          </Card>
        </div>
      </div>
    </div>
  );
};
export default ProjectSummary;
