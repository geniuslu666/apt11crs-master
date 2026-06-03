import http from '../utils/http'

export interface StaffLoginParams {
  username: string
  password: string
}

export interface StaffLoginResponse {
  token: string
  expires: number
  name: string
  mobile: string
  username: string
}

export interface VerifyLogParams {
  pageNum: number
  pageSize: number
}

export interface VerifyLogItem {
  id: number
  orderId: number
  orderSn: string
  productId: number
  memberId: number
  verifyStaffId: number
  productInfo: {
    id?: number
    title?: string
    subTitle?: string
  }
  orderInfo?: {
    id?: number
    orderSn?: string
    bookingName?: string
    phoneArea?: string
    bookingMobile?: string
    bookingNum?: number
    orderStatus?: string
  }
  bookDate: string
  verifyTime: string
}

export interface VerifyLogResponse {
  list: VerifyLogItem[]
  count: number
}

export interface VerifyLogDetailResponse {
  id: number
  orderId: number
  orderSn: string
  productId: number
  memberId: number
  bookDate: string
  verifyStaffId: number
  verifyTime: string
  staffName?: string
  productInfo: {
    id?: number
    title?: string
    subTitle?: string
  }
  orderInfo?: {
    id?: number
    orderSn?: string
    bookingName?: string
    phoneArea?: string
    bookingMobile?: string
    bookingNum?: number
  }
}

export interface CodeViewResponse {
  id: number
  orderSn: string
  productId: number
  memberId: number
  bookingName: string
  phoneArea?: string
  bookingMobile?: string
  bookingEmail?: string
  bookingNum?: number
  bookDate?: string
  orderAmount?: number
  createdAt?: string
  orderStatus?: string
  productInfo: {
    id?: number
    title?: string
    subTitle?: string
    meetingPlace?: string
    meetingTime?: string
    ggLat?: string
    ggLng?: string
  }
}

export interface CodeVerifyResponse {
  verifyTime: string
  orderSn: string
}

export interface ConfigResponse {
  contactMobile: string
}

export function staffLogin(params: StaffLoginParams) {
  return http.post<any, StaffLoginResponse>('/travelStaff/login', params)
}

export function staffLogout() {
  return http.post('/travelStaff/logout')
}

export function getVerifyLog(params: VerifyLogParams) {
  return http.post<any, VerifyLogResponse>('/travelStaff/verifyLog/list', params)
}

export function getVerifyLogView(id: number) {
  return http.post<any, VerifyLogDetailResponse>(`/travelStaff/verifyLog/view`, { id })
}

export function getCodeView(code: string) {
  return http.post<any, CodeViewResponse>('/travelStaff/code/view', { code })
}

export function verifyCode(orderSn: string) {
  return http.post<any, CodeVerifyResponse>('/travelStaff/code/verify', { orderSn })
}

export function getConfig() {
  return http.post<any, ConfigResponse>(`/travelStaff/config`)
}