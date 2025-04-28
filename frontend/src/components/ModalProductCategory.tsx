import { Button, Label, Modal, Textarea, TextInput } from "flowbite-react";
import { useState, type FC } from "react";
import {
  createProductCategory,
  updateProductCategory,
} from "../services/api/productCategoryApi";
import { ProductCategoryModel } from "../models/product";

interface ModalProductCategoryProps {
  category: ProductCategoryModel;
  setCategory: (category: ProductCategoryModel) => void;
  show: boolean;
  setShow: (show: boolean) => void;
  onCreate: () => void;
}

const ModalProductCategory: FC<ModalProductCategoryProps> = ({
  show,
  setShow,
  onCreate,
  category,
  setCategory,
}) => {
  const [name, setName] = useState("");
  return (
    <Modal show={show} onClose={() => setShow(false)}>
      <Modal.Header>Create Product Category</Modal.Header>
      <Modal.Body>
        <div className="space-y-6">
          <div className="mb-2 block">
            <Label htmlFor="name" value="Name" />
            <TextInput
              id="name"
              type="text"
              placeholder="Name"
              required={true}
              value={category?.name}
              onChange={(e) =>
                setCategory({
                  ...category!,
                  name: e.target.value,
                })
              }
              className="input-white"
            />
          </div>
          <div className="mb-2 block">
            <Label htmlFor="description" value="Description" />
            <Textarea
              id="description"
              placeholder="Description"
              rows={4}
              value={category?.description}
              onChange={(e) =>
                setCategory({
                  ...category!,
                  description: e.target.value,
                })
              }
              className="input-white"
              style={{ backgroundColor: "white" }}
            />
          </div>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <div className="justify-end flex w-full">
          <Button
            onClick={async () => {
              try {
                if (category.id) {
                  await updateProductCategory(category.id, category);
                } else {
                  await createProductCategory(category);
                }
                onCreate();
                setShow(false);
                setCategory({
                  name: "",
                  description: "",
                });
              } catch (error) {
                console.error(error);
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
export default ModalProductCategory;
