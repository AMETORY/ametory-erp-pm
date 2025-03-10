import type { FC } from "react";
import { ProjectModel } from "../models/project";
import Chart from "react-google-charts";
import { daysToMilliseconds } from "../utils/helper";

interface GanttChartProps {
  project: ProjectModel;
}

const GanttChart: FC<GanttChartProps> = ({ project }) => {
  const columns = [
    { type: "string", label: "Task ID" },
    { type: "string", label: "Task Name" },
    { type: "date", label: "Start Date" },
    { type: "date", label: "End Date" },
    { type: "number", label: "Duration" },
    { type: "number", label: "Percent Complete" },
    { type: "string", label: "Dependencies" },
  ];

  const rows = [
    [
      "Research",
      "Find sources",
      new Date(2015, 0, 1),
      new Date(2015, 0, 5),
      null,
      100,
      null,
    ],
    [
      "Write",
      "Write paper",
      null,
      new Date(2015, 0, 9),
      daysToMilliseconds(3),
      25,
      "Research,Outline",
    ],
    [
      "Cite",
      "Create bibliography",
      null,
      new Date(2015, 0, 7),
      daysToMilliseconds(1),
      20,
      "Research",
    ],
    [
      "Complete",
      "Hand in paper",
      null,
      new Date(2015, 0, 10),
      daysToMilliseconds(1),
      0,
      "Cite,Write",
    ],
    [
      "Outline",
      "Outline paper",
      null,
      new Date(2015, 0, 6),
      daysToMilliseconds(1),
      100,
      "Research",
    ],
  ];
  const data = [columns, ...rows];
  return (
    <div
      className=" px-4 flex flex-col overflow-auto"
      style={{ height: "calc(100vh - 240px)" }}
    >
      <Chart chartType="Gantt" width="100%" height="100%" data={data} />
    </div>
  );
};
export default GanttChart;
