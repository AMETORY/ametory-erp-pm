import { Button, Label, Modal, TextInput } from "flowbite-react";
import type { FC } from "react";
import { ColumnModel } from "../models/column";
import { updateColumn } from "../services/api/projectApi";
import toast from "react-hot-toast";

interface ModalColumnProps {
  projectId: string;
  column: ColumnModel;
  showModal: boolean;
  setShowModal: (val: boolean) => void;
  onChangeColumn: (column: ColumnModel) => void;
}

const ModalColumn: FC<ModalColumnProps> = ({
  projectId,
  column,
  showModal,
  setShowModal,
  onChangeColumn,
}) => {
  return (
    <Modal show={showModal} onClose={() => setShowModal(false)}>
      <Modal.Header>Edit Column</Modal.Header>
      <Modal.Body className="space-y-6">
        <div>
          <div className="mb-2 block">
            <Label htmlFor="icon" value=" Icon" />
          </div>
          <TextInput
            id="icon"
            value={column.icon}
            onChange={(e) =>
              onChangeColumn({ ...column, icon: e.target.value })
            }
            placeholder=" Icon"
          />
        </div>
        <div>
          <div className="mb-2 block">
            <Label htmlFor="name" value="Name" />
          </div>
          <TextInput
            id="name"
            value={column.name}
            onChange={(e) =>
              onChangeColumn({ ...column, name: e.target.value })
            }
            placeholder="Name"
          />
        </div>
        <div>
          <div className="mb-2 block">
            <Label htmlFor="color" value="Color" />
          </div>
          <input
            id="color"
            type="color"
            value={column.color}
            onChange={(e) =>
              onChangeColumn({ ...column, color: e.target.value })
            }
          />
        </div>
      </Modal.Body>
      <Modal.Footer className="flex justify-end">
        <Button
          onClick={() => {
            updateColumn(projectId!, {
              ...column,
            })
              .then(() => setShowModal(false))
              .catch(toast.error);
          }}
        >
          Save
        </Button>
        <Button color="gray" onClick={() => setShowModal(false)}>
          Cancel
        </Button>
      </Modal.Footer>
    </Modal>
  );
};
export default ModalColumn;
