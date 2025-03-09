import { useContext, useRef, useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { ProfileContext } from "../contexts/ProfileContext";
import {
  Button,
  FileInput,
  Tabs,
  TabsRef,
  Textarea,
  TextInput,
} from "flowbite-react";
import { uploadFile } from "../services/api/commonApi";
import { FileModel } from "../models/file";
import { LoadingContext } from "../contexts/LoadingContext";
import toast from "react-hot-toast";
import { changePassword, updateProfile } from "../services/api/authApi";
import { BsInfoCircle, BsLock } from "react-icons/bs";
import { GoLock } from "react-icons/go";

interface ProfilePageProps {}

const ProfilePage: FC<ProfilePageProps> = ({}) => {
  const { loading, setLoading } = useContext(LoadingContext);
  const { profile, setProfile } = useContext(ProfileContext);
  const [file, setFile] = useState<FileModel>();
  const tabsRef = useRef<TabsRef>(null);
  const [activeTab, setActiveTab] = useState(0);

  const renderInfo = () => (
    <div className="flex flex-col gap-4">
      <h1 className="text-3xl font-bold">Edit Profile</h1>
      <div className="bg-white rounded-lg p-4">
        <div className="flex flex-col gap-2 space-y-4">
          {profile?.profile_picture && (
            <div className="flex justify-center py-4 items-center">
              <img
                className="w-64 h-64 aspect-square object-cover rounded-full"
                src={profile?.profile_picture?.url}
                alt="profile"
              />
            </div>
          )}

          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Profile Picture</label>
            <FileInput
              accept="image/*"
              id="file-upload"
              onChange={(el) => {
                if (el.target.files) {
                  let f = el.target.files[0];
                  if (!f) return;
                  uploadFile(f, {}, (val) => {
                    console.log(val);
                  }).then((v: any) => {
                    setFile(v.data);
                    setProfile({
                      ...profile!,
                      profile_picture: v.data,
                    });
                  });
                }
              }}
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Name</label>
            <TextInput
              type="text"
              value={profile?.full_name}
              onChange={(e) =>
                setProfile({ ...profile!, full_name: e.target.value })
              }
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Address</label>
            <Textarea
              value={profile?.address}
              onChange={(e) =>
                setProfile({ ...profile!, address: e.target.value })
              }
            />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Email</label>
            <TextInput
              readOnly
              type="email"
              value={profile?.email}
              onChange={(e) =>
                setProfile({ ...profile!, email: e.target.value })
              }
            />
          </div>
          <div>
            <Button
              type="submit"
              className="mt-8 w-32"
              onClick={async () => {
                try {
                  setLoading(true);
                  await updateProfile(profile!);
                  toast.success("Profile updated successfully");
                } catch (error) {
                  toast.error(`${error}`);
                } finally {
                  setLoading(false);
                }
              }}
            >
              Save
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
  const renderSecurity = () => (
    <form
      className="flex flex-col gap-4"
      onSubmit={async (e) => {
        e.preventDefault();
        try {
          setLoading(true);
          const { old_password, new_password, confirm_password } =
            Object.fromEntries(new FormData(e.target as HTMLFormElement));
          if (new_password !== confirm_password) {
            toast.error("Password and confirm password does not match");
            return;
          }
          await changePassword({
            old_password: old_password as string,
            new_password: new_password as string,
          });
          toast.success("Password changed successfully");
        } catch (error) {
          toast.error(`${error}`);
        } finally {
          setLoading(false);
        }
      }}
    >
      <div className="flex flex-col gap-1">
        <label className="text-sm font-medium">Old Password</label>
        <TextInput
          type="password"
          name="old_password"
          placeholder="Input your old password"
          required
        />
      </div>
      <div className="flex flex-col gap-1">
        <label className="text-sm font-medium">New Password</label>
        <TextInput
          type="password"
          name="new_password"
          placeholder="Input your new password"
          required
        />
      </div>
      <div className="flex flex-col gap-1">
        <label className="text-sm font-medium">Confirm Password</label>
        <TextInput
          type="password"
          name="confirm_password"
          placeholder="Input your confirm password"
          required
        />
      </div>
      <div>
        <Button type="submit" className="mt-8">
          Change Password
        </Button>
      </div>
    </form>
  );
  return (
    <AdminLayout>
      <div className="w-full h-full flex flex-col gap-4 px-8">
        <Tabs
          aria-label="Default tabs"
          variant="default"
          ref={tabsRef}
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
          className="mt-4"
        >
          <Tabs.Item
            active={activeTab === 0}
            title="Basic Info"
            icon={BsInfoCircle}
          >
            {renderInfo()}
          </Tabs.Item>
          <Tabs.Item
            active={activeTab === 1}
            title="Security"
            icon={GoLock}
          >
           {renderSecurity()}
          </Tabs.Item>
        </Tabs>
      </div>
    </AdminLayout>
  );
};
export default ProfilePage;
