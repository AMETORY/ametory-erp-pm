import { Button, Pagination, Table } from "flowbite-react";
import { useContext, useEffect, useState, type FC } from "react";
import { LuFilter } from "react-icons/lu";
import ModalProductCategory from "./ModalProductCategory";
import { LoadingContext } from "../contexts/LoadingContext";
import { PaginationResponse } from "../objects/pagination";
import {
  deleteProductCategory,
  getProductCategories,
} from "../services/api/productCategoryApi";
import { SearchContext } from "../contexts/SearchContext";
import { getPagination } from "../utils/helper";
import { ProductCategoryModel } from "../models/product";

interface ProductCategoryTableProps {}

const ProductCategoryTable: FC<ProductCategoryTableProps> = ({}) => {
  const [mounted, setMounted] = useState(false);
  const [showModal, setShowModal] = useState(false);
  const { loading, setLoading } = useContext(LoadingContext);
  const { search, setSearch } = useContext(SearchContext);
  const [productCategories, setProductCategories] = useState<
    ProductCategoryModel[]
  >([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [pagination, setPagination] = useState<PaginationResponse>();
  const [category, setCategory] = useState<ProductCategoryModel>();

  useEffect(() => {
    setMounted(true);
  }, []);
  useEffect(() => {
    if (mounted) {
      getAllCategories();
    }
  }, [mounted, page, size, search]);

  const getAllCategories = () => {
    getProductCategories({ page, size, search }).then((res: any) => {
      setProductCategories(res.data.items);
      setPagination(getPagination(res.data));
    });
  };
  return (
    <div>
      <div className="flex justify-between items-center mb-4">
        <h1 className="text-3xl font-bold ">Category</h1>
        <div className="flex items-center gap-2">
          <Button
            gradientDuoTone="purpleToBlue"
            pill
            onClick={() => {
                setCategory({
                    name: "",
                    description: "",
                });
              setShowModal(true);
            }}
          >
            + Category
          </Button>
          <LuFilter
            className=" cursor-pointer text-gray-400 hover:text-gray-600"
            onClick={() => {}}
          />
        </div>
      </div>
      <Table hoverable>
        <Table.Head>
          <Table.HeadCell>Name</Table.HeadCell>
          <Table.HeadCell>Description</Table.HeadCell>
          <Table.HeadCell></Table.HeadCell>
        </Table.Head>
        <Table.Body>
          {productCategories.length === 0 && (
            <Table.Row>
              <Table.Cell colSpan={5} className="text-center">
                No data found.
              </Table.Cell>
            </Table.Row>
          )}
          {productCategories.map((category) => (
            <Table.Row key={category.id}>
              <Table.Cell>{category.name}</Table.Cell>
              <Table.Cell>{category.description}</Table.Cell>
              <Table.Cell>
                <a
                  
                  className="font-medium text-cyan-600 hover:underline dark:text-cyan-500 cursor-pointer"
                  onClick={() => {
                    setCategory(category)
                    setShowModal(true);
                  }}
                >
                  Edit
                </a>
                <a
                  className="font-medium text-red-600 hover:underline dark:text-red-500 ms-2 cursor-pointer"
                  onClick={(e) => {
                    e.preventDefault();
                    if (
                      window.confirm(
                        `Are you sure you want to delete  ${category.name}?`
                      )
                    ) {
                      deleteProductCategory(category?.id!).then(() => {
                        getAllCategories();
                      });
                    }
                  }}
                >
                  Delete
                </a>
              </Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
      <Pagination
        className="mt-4"
        currentPage={page}
        totalPages={pagination?.total_pages ?? 0}
        onPageChange={(val) => {
          setPage(val);
        }}
        showIcons
      />
      {category && (
        <ModalProductCategory
          show={showModal}
          setShow={setShowModal}
          onCreate={getAllCategories}
          category={category!}
          setCategory={setCategory}
        />
      )}
    </div>
  );
};
export default ProductCategoryTable;
