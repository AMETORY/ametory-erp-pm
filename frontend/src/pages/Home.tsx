import { Button, Datepicker, Modal } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_TOKEN } from "../utils/constants";
import AdminLayout from "../components/layouts/admin";
import Chart from "react-google-charts";
import { MemberModel } from "../models/member";
import { getMembers } from "../services/api/commonApi";
import toast from "react-hot-toast";
import moment from "moment";
import Select from "react-select";

interface HomeProps {}

const Home: FC<HomeProps> = ({}) => {
  const [members, setMembers] = useState<MemberModel[]>([]);
  const [selectedMembers, setSelectedMembers] = useState<
    { value: string; label: string }[]
  >([]);
  const [mounted, setMounted] = useState(false);
  const today = new Date();
  const start = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate()
  );
  const end = new Date(
    today.getFullYear(),
    today.getMonth(),
    today.getDate(),
    23,
    59,
    59
  );
  const [dateRange, setDateRange] = useState([start, end]);
  const [modalDateOpen, setModalDateOpen] = useState(false);

  useEffect(() => {
    setMounted(true);

    return () => {
      setMounted(false);
    };
  }, []);

  useEffect(() => {
    if (mounted) {
      getMembers({ page: 1, size: 10 })
        .then((res: any) => {
          setMembers(res.data.items);
        })
        .catch(toast.error);
    }
  }, [mounted]);

  return (
    <AdminLayout>
      <div className="p-4 h-[calc(100vh-100px)] overflow-y-auto">
        <div className="flex flex-row p-2 rounded-lg bg-gray-100 min-h-[60px] justify-between items-center">
          <div className="p-2 bg-white rounded-lg min-w-[320px]">
            <Select
              className="w-full"
              isMulti
              placeholder="Select Members"
              options={members.map((member) => ({
                label: member.user?.full_name ?? "",
                value: member.id ?? "",
              }))}
              value={selectedMembers}
              onChange={(selectedOptions) => {
                setSelectedMembers(selectedOptions.map((e) => e));
                // setSelectedMembers(selectedOptions.map((option) => ({ id: option.value })));
              }}
            />
          </div>
          <div
            className="p-2 bg-white rounded-lg min-w-[240px] cursor-pointer"
            onClick={() => setModalDateOpen(true)}
          >
            {moment(dateRange[0]).format("DD MMM YYYY")}{" "}
            {moment(dateRange[0]).format("HH:mm")} -{" "}
            {moment(dateRange[1]).format("DD MMM YYYY")}{" "}
            {moment(dateRange[1]).format("HH:mm")}
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div className="mt-4 border rounded-lg p-2 hover:shadow-lg cursor-pointer">
            <Chart
              width="100%"
              height="300px"
              chartType="PieChart"
              loader={<div>Loading Chart</div>}
              data={[
                ["Conversation", "Total"],
                ["New Customer", Math.floor(Math.random() * 100) + 1],
                ["Old Customer", Math.floor(Math.random() * 100) + 1],
              ]}
              options={{
                title: "Customer Interaction",
                pieHole: 0.4,
              }}
              rootProps={{ "data-testid": "7" }}
            />
          </div>
          <div className="mt-4 border rounded-lg p-2 hover:shadow-lg cursor-pointer">
            <Chart
              width="100%"
              height="300px"
              chartType="PieChart"
              loader={<div>Loading Chart</div>}
              data={[
                ["Conversation", "Seconds"],
                ["ATR old Customer", Math.floor(Math.random() * 60) + 1],
                ["ATR new Customer", Math.floor(Math.random() * 60) + 1],
              ]}
              options={{
                title: "ATR Customer Interaction",
                legend: { position: "bottom" },
                is3D: true,
              }}
              rootProps={{ "data-testid": "7" }}
            />
          </div>
        </div>
        <div className="grid grid-cols-2 gap-4">
          <div>
            <div className="mt-4 border rounded-lg p-2 hover:shadow-lg cursor-pointer">
              <Chart
                width="100%"
                height="300px"
                chartType="ColumnChart"
                loader={<div>Loading Chart</div>}
                data={[
                  ["Hour", "New Customer", "Old Customer"],
                  ...Array.from({ length: 24 }).map((_, i) => [
                    `${i < 10 ? "0" + i : i}:00`,
                    Math.floor(Math.random() * 100) + 1,
                    Math.floor(Math.random() * 100) + 1,
                  ]),
                ]}
                options={{
                  title: "Total Inbound per Hour",
                  legend: { position: "bottom" },
                  hAxis: { title: "Hour" },
                  vAxis: { title: "Number of Inbound" },
                }}
                rootProps={{ "data-testid": "1" }}
              />
            </div>
          </div>
          <div>
            <div className="mt-4 border rounded-lg p-2 hover:shadow-lg cursor-pointer">
              <Chart
                width="100%"
                height="300px"
                chartType="LineChart"
                loader={<div>Loading Chart</div>}
                data={[
                  ["Hour", "ATR New Customer", "ATR Old Customer"],
                  ...Array.from({ length: 24 }).map((_, i) => [
                    `${i < 10 ? "0" + i : i}:00`,
                    Math.floor(Math.random() * 100) +
                      Math.floor(Math.random() * 10) +
                      1,
                    Math.floor(Math.random() * 100) +
                      Math.floor(Math.random() * 10) +
                      1,
                  ]),
                ]}
                options={{
                  title: "ATR per Hour",
                  legend: { position: "bottom" },
                  hAxis: { title: "Hour" },
                  vAxis: { title: "Second" },
                  curveType: "function",
                }}
                rootProps={{ "data-testid": "2" }}
              />
            </div>
          </div>
        </div>
      </div>
      <Modal
        size="4xl"
        show={modalDateOpen}
        onClose={() => setModalDateOpen(false)}
        dismissible
      >
        <Modal.Header>Date Range</Modal.Header>
        <Modal.Body>
          <div className="flex flex-col pb-32">
            <div className="grid grid-cols-2 gap-2 ">
              <div className="flex daterange">
                <Datepicker
                  value={dateRange[0]}
                  onChange={(v) => setDateRange([v!, dateRange[1]])}
                  className="min-w-[200px]"
                />
                <div className="flex w-full">
                  <input
                    type="time"
                    id="time"
                    className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={moment(dateRange[0]).format("HH:mm")}
                    onChange={(e) => {
                      const newDate = new Date(dateRange[0]);
                      newDate.setHours(parseInt(e.target.value.split(":")[0]));
                      newDate.setMinutes(
                        parseInt(e.target.value.split(":")[1])
                      );
                      setDateRange([newDate, dateRange[1]]);
                    }}
                  />
                  <span className="inline-flex  items-center px-3 text-sm text-gray-900 bg-gray-200  rounded-s-0 border-0 border-gray-300 rounded-e-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <svg
                      className="w-4 h-4 text-gray-500 dark:text-gray-400"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fillRule="evenodd"
                        d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586V8Z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </span>
                </div>
              </div>
              <div className="flex daterange">
                <Datepicker
                  value={dateRange[1]}
                  onChange={(v) => setDateRange([dateRange[0], v!])}
                  className="min-w-[200px]"
                />
                <div className="flex w-full">
                  <input
                    type="time"
                    id="time"
                    className="rounded-none rounded-s-lg bg-gray-50 border text-gray-900 leading-none focus:ring-blue-500 focus:border-blue-500 block flex-1 w-full text-sm border-gray-300 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
                    value={moment(dateRange[1]).format("HH:mm")}
                    onChange={(e) => {
                      const newDate = new Date(dateRange[1]);
                      newDate.setHours(parseInt(e.target.value.split(":")[0]));
                      newDate.setMinutes(
                        parseInt(e.target.value.split(":")[1])
                      );
                      setDateRange([dateRange[0], newDate]);
                    }}
                  />
                  <span className="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200  rounded-s-0 border-0 border-gray-300 rounded-e-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
                    <svg
                      className="w-4 h-4 text-gray-500 dark:text-gray-400"
                      aria-hidden="true"
                      xmlns="http://www.w3.org/2000/svg"
                      fill="currentColor"
                      viewBox="0 0 24 24"
                    >
                      <path
                        fillRule="evenodd"
                        d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10-4.477 10-10 10S2 17.523 2 12Zm11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586V8Z"
                        clipRule="evenodd"
                      />
                    </svg>
                  </span>
                </div>
              </div>
            </div>
            <div className="mt-4">
              <ul className="grid grid-cols-2 gap-4">
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const start = new Date();
                      start.setHours(0, 0, 0, 0);
                      const end = new Date();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([start, end]);
                    }}
                  >
                    Today
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const yesterday = new Date();
                      yesterday.setDate(yesterday.getDate() - 1);
                      yesterday.setHours(23, 59, 59, 999);
                      setDateRange([yesterday, yesterday]);
                    }}
                  >
                    Yesterday
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const lastWeek = new Date();
                      lastWeek.setDate(lastWeek.getDate() - 7);
                      lastWeek.setHours(0, 0, 0, 0);
                      const end = new Date();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([lastWeek, end]);
                    }}
                  >
                    Last Week
                  </button>
                </li>
                <li>
                  <button
                    className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                    onClick={() => {
                      const firstDayOfMonth = new Date(
                        new Date().getFullYear(),
                        new Date().getMonth(),
                        1
                      );
                      firstDayOfMonth.setHours(0, 0, 0, 0);
                      const end = moment(firstDayOfMonth).endOf("month").toDate();
                      end.setHours(23, 59, 59, 999);
                      setDateRange([firstDayOfMonth, end]);
                    }}
                  >
                    This Month
                  </button>
                </li>
                {[...Array(4)].map((_, i) => {
                  const quarter = i + 1;
                  const start = new Date(
                    new Date().getFullYear(),
                    (quarter - 1) * 3,
                    1
                  );
                  const end = new Date(
                    new Date().getFullYear(),
                    quarter * 3,
                    0
                  );
                  return (
                    <li key={i}>
                      <button
                        key={i}
                        className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                        onClick={() => {
                          const endCopy = new Date(end);
                          endCopy.setHours(23, 59, 59, 999);
                          setDateRange([start, endCopy]);
                        }}
                      >
                        Q{quarter}
                      </button>
                    </li>
                  );
                })}

                {[...Array(4)].map((_, i) => {
                  const year = new Date().getFullYear() - (i === 0 ? 0 : i);
                  return (
                    <li key={i}>
                      <button
                        key={i}
                        className="rs-input clear-start text-center hover:font-semibold hover:bg-gray-50"
                        onClick={() => {
                          const endCopy = new Date(year, 11, 31);
                          endCopy.setHours(23, 59, 59, 999);
                          setDateRange([new Date(year, 0, 1), endCopy]);
                        }}
                      >
                        {year == new Date().getFullYear() ? "This Year" : year}
                      </button>
                    </li>
                  );
                })}
              </ul>
            </div>
          </div>
        </Modal.Body>
      </Modal>
    </AdminLayout>
  );
};
export default Home;
