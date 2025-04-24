import { useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Button } from "flowbite-react";

interface BroadcastPageProps {}

const BroadcastPage: FC<BroadcastPageProps> = ({}) => {
  const [showModal, setShowModal] = useState(false);
  return (
    <AdminLayout>
      <div className="p-8">
        <div className="flex justify-between items-center mb-4">
          <h1 className="text-3xl font-bold ">Broadcast</h1>
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
              setShowModal(true);
            }}
          >
            + Create new broadcast
          </Button>
        </div>
      </div>
    </AdminLayout>
  );
};
export default BroadcastPage;
