import { CompanyModel } from "./company";
import { FileModel } from "./file";
import { TagModel } from "./tag";

export interface ProductModel {
  id?: string;
  name?: string;
  description?: string;
  display_name?: string;
  sku?: string;
  barcode?: string;
  price?: number;
  original_price?: number;
  company_id?: string;
  company?: CompanyModel;
  category_id?: string;
  category?: ProductCategoryModel;
  product_images?: FileModel[];
  total_stock?: number;
  tags?: TagModel[];
  height?: number;
  length?: number;
  weight?: number;
  width?: number;
  status?: string;
}

export interface ProductCategoryModel {
  id?: string;
  name: string;
  description?: string;
  color?: string;
  icon_url?: string;
}
