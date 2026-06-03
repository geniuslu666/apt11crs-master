import { defineStore } from 'pinia'

const TOKEN_KEY = 'H5_STAFF_TOKEN'
const INFO_KEY = 'H5_STAFF_INFO'

export interface StaffInfo {
  username: string
  name: string
  mobile: string
}

export const useTravelStaffStore = defineStore('travelStaff', {
  state: () => ({
    token: localStorage.getItem(TOKEN_KEY) || '',
    staffInfo: JSON.parse(localStorage.getItem(INFO_KEY) || 'null') as StaffInfo | null,
  }),
  getters: {
    isLoggedIn: (state) => !!state.token,
  },
  actions: {
    setToken(token: string) {
      this.token = token
      localStorage.setItem(TOKEN_KEY, token)
    },
    setStaffInfo(info: StaffInfo) {
      this.staffInfo = info
      localStorage.setItem(INFO_KEY, JSON.stringify(info))
    },
    clear() {
      this.token = ''
      this.staffInfo = null
      localStorage.removeItem(TOKEN_KEY)
      localStorage.removeItem(INFO_KEY)
    },
  },
})
