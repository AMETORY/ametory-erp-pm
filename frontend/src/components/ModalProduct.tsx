import {
  Button,
  FileInput,
  Label,
  Modal,
  Textarea,
  TextInput,
} from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { ProductCategoryModel, ProductModel } from "../models/product";
import {
  createProduct,
  deleteProductImage,
  updateProduct,
} from "../services/api/productApi";
import toast from "react-hot-toast";
import { getProductCategories } from "../services/api/productCategoryApi";
import Select, { InputActionMeta } from "react-select";

import CurrencyInput from "react-currency-input-field";
import Barcode from "react-barcode";
import { uploadFile } from "../services/api/commonApi";
interface ModalProductProps {
  show: boolean;
  setShow: (show: boolean) => void;
  product?: ProductModel | undefined;
  setProduct: (product: ProductModel) => void;
  onCreateProduct: (product: ProductModel) => void;
}

const ModalProduct: FC<ModalProductProps> = ({
  show,
  setShow,
  product,
  setProduct,
  onCreateProduct,
}) => {
  const [categories, setCategories] = useState<ProductCategoryModel[]>([]);
  const handleCreateProduct = async () => {
    try {
      if (product?.id) {
        const res: any = await updateProduct(product!.id, product);
        onCreateProduct(res.data);
      } else {
        const res: any = await createProduct(product);
        onCreateProduct(res.data);
      }

      setShow(false);
    } catch (error) {
      toast.error(`${error}`);
    }
  };

  useEffect(() => {
    searchCategory("");
  }, []);

  const searchCategory = (s: string) => {
    getProductCategories({ page: 1, size: 10, search: s }).then((res: any) => {
      setCategories(res.data.items);
    });
  };
  return (
    <Modal show={show} onClose={() => setShow(false)}>
      <Modal.Header>Create Product</Modal.Header>
      <Modal.Body>
        <div className="flex flex-col space-y-4">
          {(product?.product_images ?? []).length > 0 && (
            <div className="flex justify-center py-4 items-center">
              <img
                className="w-64 h-64 aspect-square object-cover rounded-full"
                src={product?.product_images![0].url}
                alt="profile"
              />
            </div>
          )}

          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium">Image</label>
            <FileInput
              accept="image/*"
              id="file-upload"
              onChange={async (el) => {
                if (el.target.files) {
                  let f = el.target.files[0];
                  if (!f) return;

                  if ((product?.product_images ?? []).length > 0) {
                    await deleteProductImage(
                      product!.id!,
                      product!.product_images![0].id!
                    );
                  }
                  uploadFile(
                    f,
                    {
                      ref_id: product!.id!,
                      ref_type: "product",
                    },
                    (val) => {
                      console.log(val);
                    }
                  ).then((v: any) => {
                    setProduct({
                      ...product!,
                      product_images: [ v.data],
                    });
                  });
                }
              }}
            />
          </div>
          <div className="mb-4">
            <Label htmlFor="product-name" value="Product Name" />
            <TextInput
              id="product-name"
              placeholder="Product Name"
              value={product?.name ?? ""}
              onChange={(e) =>
                setProduct({ ...product!, name: e.target.value })
              }
              className="input-white"
            />
          </div>
          <div className="mb-4">
            <Label htmlFor="product-sku" value="SKU" />
            <TextInput
              id="product-sku"
              placeholder="Product SKU"
              value={product?.sku ?? ""}
              onChange={(e) => setProduct({ ...product!, sku: e.target.value })}
              className="input-white"
            />
          </div>
          <div className="mb-4">
            <Label htmlFor="product-barcode" value="Barcode" />
            <TextInput
              id="product-barcode"
              placeholder="Product Barcode"
              value={product?.barcode ?? ""}
              onChange={(e) =>
                setProduct({ ...product!, barcode: e.target.value })
              }
              className="input-white"
            />
            {product?.barcode && (
              <Barcode className="mt-2" height={50} value={product.barcode} />
            )}
          </div>
          <div className="mb-4">
            <Label htmlFor="product-price" value="Product Price" />
            <CurrencyInput
              className="rs-input !p-1.5 "
              value={product?.price ?? 0}
              groupSeparator="."
              decimalSeparator=","
              onValueChange={(value, name, values) => {
                setProduct({
                  ...product!,
                  price: values?.float ?? 0,
                });
              }}
            />
          </div>
          <div className="mb-4">
            <Label htmlFor="product-category" value="Category" />
            <Select
              id="product-category"
              value={
                product?.category
                  ? { label: product.category.name, value: product.category.id }
                  : null
              }
              onChange={(e) =>
                setProduct({
                  ...product!,
                  category: { id: e!.value, name: e!.label },
                  category_id: e!.value,
                })
              }
              options={categories.map((c) => ({ label: c.name, value: c.id }))}
              onInputChange={(e) => searchCategory(e)}
            />
          </div>
          <div className="mb-4">
            <Label htmlFor="product-description" value="Description" />
            <Textarea
              rows={7}
              id="product-description"
              placeholder="Product Description"
              value={product?.description ?? ""}
              onChange={(e) =>
                setProduct({ ...product!, description: e.target.value })
              }
              className="input-white"
              style={{ backgroundColor: "white" }}
            />
          </div>
        </div>
      </Modal.Body>
      <Modal.Footer>
        <div className="flex justify-end w-full">
          <Button onClick={handleCreateProduct}>Save</Button>
        </div>
      </Modal.Footer>
    </Modal>
  );
};
export default ModalProduct;
