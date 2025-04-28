import { useState, type FC } from "react";
import AdminLayout from "../components/layouts/admin";
import { Tabs } from "flowbite-react";
import { RiShoppingBagLine } from "react-icons/ri";
import { BsListCheck } from "react-icons/bs";
import ProductTable from "../components/ProductTable";
import ProductCategoryTable from "../components/ProductCategoryTable";

interface ProductPageProps {}

const ProductPage: FC<ProductPageProps> = ({}) => {
  const [activeTab, setActiveTab] = useState(0);
  return (
    <AdminLayout>
      <div className="w-full flex flex-col gap-4 px-8">
        <Tabs
          aria-label="Default tabs"
          variant="default"
          onActiveTabChange={(tab) => {
            setActiveTab(tab);
            // console.log(tab);
          }}
          className="mt-4"
        >
          <Tabs.Item
            active={activeTab === 0}
            title="Product"
            icon={RiShoppingBagLine}
          >
            <ProductTable />
          </Tabs.Item>
          <Tabs.Item
            active={activeTab === 1}
            title="Category"
            icon={BsListCheck}
          >
            <ProductCategoryTable />
          </Tabs.Item>
        </Tabs>
      </div>
    </AdminLayout>
  );
};
export default ProductPage;
