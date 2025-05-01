import { Modal, Table } from "flowbite-react";
import { useEffect, useState, type FC } from "react";
import { HiMagnifyingGlass } from "react-icons/hi2";
import { ProductModel } from "../models/product";
import { PaginationResponse } from "../objects/pagination";
import { getProducts } from "../services/api/productApi";
import { getPagination, money } from "../utils/helper";

interface ModalProductListProps {
  show: boolean;
  setShow: (show: boolean) => void;
  selectProduct: (product: ProductModel) => void;
}

const ModalProductList: FC<ModalProductListProps> = ({ show, setShow, selectProduct }) => {
  const [search, setSearch] = useState("");
  const [products, setProducts] = useState<ProductModel[]>([]);
  const [page, setPage] = useState(1);
  const [size, setsize] = useState(20);
  const [pagination, setPagination] = useState<PaginationResponse>();

  useEffect(() => {
    getAllProducts();
  }, [page, size, search]);

  const getAllProducts = () => {
    getProducts({ page, size, search }).then((res: any) => {
      setProducts(res.data.items);
      setPagination(getPagination(res.data));
    });
  };

  const searchBox = (
    <div className="relative w-full mb-4 mr-6 focus-within:text-purple-500">
      <div className="absolute inset-y-0 left-0 flex items-center pl-3">
        <HiMagnifyingGlass />
      </div>
      <input
        type="text"
        className="w-full py-2 pl-10 text-sm text-gray-700 bg-white border border-gray-300 rounded-2xl shadow-sm focus:outline-none focus:ring focus:ring-indigo-200 focus:border-indigo-500"
        placeholder="Search"
        value={search}
        onChange={(e) => setSearch(e.target.value)}
      />
    </div>
  );

  return (
    <Modal show={show} onClose={() => setShow(false)}>
      <Modal.Header>List Product</Modal.Header>
      <Modal.Body>
        <div className="space-y-6 flex flex-col">{searchBox}</div>
        <div className="overflow-x-auto">
          <Table striped>
            <Table.Head>
              <Table.HeadCell>Name</Table.HeadCell>
              <Table.HeadCell>SKU</Table.HeadCell>
              <Table.HeadCell>Price</Table.HeadCell>
            </Table.Head>
            <Table.Body>
              {products.length === 0 && (
                <Table.Row>
                  <Table.Cell colSpan={5} className="text-center">
                    No data found.
                  </Table.Cell>
                </Table.Row>
              )}
              {products.map((product) => (
                <Table.Row
                  key={product.id}
                  className="bg-white dark:border-gray-700 dark:bg-gray-800 hover:bg-gray-50  cursor-pointer"
                onClick={() => {
                    selectProduct(product);
                }} >
                  <Table.Cell>{product.name}</Table.Cell>
                  <Table.Cell>{product.sku}</Table.Cell>
                  <Table.Cell>{money(product.price)}</Table.Cell>
                </Table.Row>
              ))}
            </Table.Body>
          </Table>
        </div>
      </Modal.Body>
    </Modal>
  );
};
export default ModalProductList;
