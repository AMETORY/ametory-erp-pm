import { Button, Label, Modal, TextInput } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { WhatsappMessageSessionModel } from "../models/whatsapp_message";
import { ConnectionModel } from "../models/connection";
import { getConnections } from "../services/api/connectionApi";
import Select from "react-select";
import CreatableSelect from "react-select/creatable";
import { createTag, getTags } from "../services/api/tagApi";
import { getContrastColor, randomColor } from "../utils/helper";
import { TagModel } from "../models/tag";
import toast from "react-hot-toast";
import { ContactModel } from "../models/contact";
import { updateContact } from "../services/api/contactApi";

interface ModalSessionProps {
  show: boolean;
  onClose: () => void;
  onSave: (val: WhatsappMessageSessionModel) => void;
  session?: WhatsappMessageSessionModel;
}

const ModalSession: FC<ModalSessionProps> = ({
  show,
  onClose,
  session,
  onSave,
}) => {
  const [connections, setConnections] = useState<ConnectionModel[]>([]);
  const [tags, setTags] = useState<TagModel[]>([]);
  const [selectedContact, setSelectedContact] = useState<ContactModel>();

  useEffect(() => {
    if (!session) return;
    getConnections({ page: 1, size: 50 }).then((resp: any) => {
      setConnections(resp.data);
    });
    getAllTags();
    setSelectedContact(session?.contact);
  }, [session]);

  const getAllTags = async () => {
    try {
      let resp: any = await getTags({ page: 1, size: 100 });
      setTags(resp.data.items);
    } catch (error) {
      toast.error(`${error}`);
    } finally {
    }
  };
  if (!session) return null;
  return (
    <Modal show={show} onClose={onClose}>
      <Modal.Header>{session?.contact?.name}</Modal.Header>
      <Modal.Body>
        <div className="flex flex-col space-y-4 pb-32">
          <div>
            <Label>Contact Name</Label>
            <TextInput
              type="text"
              value={selectedContact?.name}
              onChange={(e) =>
                setSelectedContact({
                  ...selectedContact!,
                  name: e.target.value,
                })
              }
            />
          </div>

          <div>
            <Label htmlFor="name" value="Connection" />
            <Select
              options={connections.map((c) => ({
                value: c.id,
                label: c.name,
              }))}
              value={{
                value: session?.ref?.id,
                label: session?.ref?.name,
              }}
              onChange={(e) => {
                onSave({
                  ...session!,
                  ref: {
                    ...session?.ref!,
                    id: e?.value,
                    name: e?.label,
                  },
                });
              }}
            />
          </div>
          <div>
            <Label htmlFor="tag" value="Tag" />
            <CreatableSelect
              id="tag"
              name="tag"
              onCreateOption={(e) => {
                console.log(e);
                createTag({
                  name: e,
                  color: randomColor({ luminosity: "dark" }),
                }).then(() => {
                  getAllTags();
                });
              }}
              isMulti={true}
              options={tags.map((tag) => ({
                value: tag.id,
                label: tag.name,
                color: tag.color,
              }))}
              value={(selectedContact?.tags ?? []).map((tag) => ({
                value: tag.id,
                label: tag.name,
                color: tag.color,
              }))}
              onChange={(e) => {
                setSelectedContact({
                  ...selectedContact!,
                  tags: e.map((tag) => ({
                    id: tag.value,
                    name: tag.label,
                    color: tag.color,
                  })),
                });
              }}
              formatOptionLabel={(option) => (
                <div
                  className="w-fit px-2 py-1 rounded-lg"
                  style={{
                    backgroundColor: option.color,
                    color: getContrastColor(option.color),
                  }}
                >
                  <span>{option.label}</span>
                </div>
              )}
              formatGroupLabel={(option) => (
                <div
                  className="w-fit px-2 py-1 rounded-lg"
                  style={{ backgroundColor: "white" }}
                >
                  <span>{option.label}</span>
                </div>
              )}
            />
          </div>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <div className="flex w-full justify-end">
          <Button
            onClick={async () => {
              if (session) {
                session.contact = selectedContact;
                onSave(session);
              }
            }}
          >
            Save
          </Button>
        </div>
      </Modal.Footer>
    </Modal>
  );
};
export default ModalSession;
