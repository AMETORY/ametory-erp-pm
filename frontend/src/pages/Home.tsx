import { Button } from "flowbite-react";
import type { FC } from "react";
import { asyncStorage } from "../utils/async_storage";
import { LOCAL_STORAGE_TOKEN } from "../utils/constants";
import AdminLayout from "../components/layouts/admin";

interface HomeProps {}

const Home: FC<HomeProps> = ({}) => {
  return (
    <AdminLayout>
      <div className="p-4">
        <h1>Home</h1>

        <Button
          gradientDuoTone="purpleToBlue"
          pill
          onClick={async () => {
            await asyncStorage.removeItem(LOCAL_STORAGE_TOKEN);
            window.location.reload();
          }}
        >
          Logout
        </Button>
      </div>
    </AdminLayout>
  );
};
export default Home;
