import { Button, Tabs } from "flowbite-react";
import { useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { SiGoogleforms } from "react-icons/si";
import { FaWpforms } from "react-icons/fa6";

interface FormPageProps {}

const FormPage: FC<FormPageProps> = ({}) => {
  const [showModal, setShowModal] = useState(false);
  const [showModalForm, setShowModalForm] = useState(false);
  const [activeTab, setActiveTab] = useState(0);
  return (
    <AdminLayout>
      <div className="p-4">
        <Tabs
          aria-label="Default tabs"
          variant="default"
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
        >
          <Tabs.Item
            title="Form Template"
            active={activeTab == 0}
            icon={SiGoogleforms}
          >
            <div className="p-4">
              <div className="flex justify-between items-center mb-4">
                <h1 className="text-3xl font-bold ">Form Template</h1>
                <Button
                  gradientDuoTone="purpleToBlue"
                  pill
                  onClick={() => {
                    setShowModal(true);
                  }}
                >
                  + Create Form Template
                </Button>
              </div>
            </div>
          </Tabs.Item>
          <Tabs.Item title="Form" active={activeTab == 1} icon={FaWpforms}>
            <div className="p-4">
              <div className="flex justify-between items-center mb-4">
                <h1 className="text-3xl font-bold ">Form </h1>
                <Button
                  gradientDuoTone="purpleToBlue"
                  pill
                  onClick={() => {
                    setShowModalForm(true);
                  }}
                >
                  + Create Form
                </Button>
              </div>
            </div>
          </Tabs.Item>
        </Tabs>
      </div>
    </AdminLayout>
  );
};
export default FormPage;
