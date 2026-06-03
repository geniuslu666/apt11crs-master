import { EmployeeDepartmentInter } from './employeeDepartment';

export interface EmployeeInter {
  id: number;
  name: string;
  phoneArea: string;
  phone: string;
  departmentDetail: EmployeeDepartmentInter;
}