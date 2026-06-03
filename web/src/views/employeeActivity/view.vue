<template>
    <div>
        <n-drawer v-model:show="showModal" :width="dialogWidth">
            <n-drawer-content closable :header-style="{
                                padding: '20px',
                              }" :body-content-style="{
                                padding: '25px 0 0',
                              }">
                <template #header>
                    <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">活动详情</div>
                </template>
                <n-spin :show="loading" description="请稍候...">
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            活动名称
                        </div>
                        <n-grid y-gap="25" :cols="3">
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动名称-中文</div>
                                <div class="level-detail-div-item-content">{{ info.zh_name ? info.zh_name : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动名称-日本语</div>
                                <div class="level-detail-div-item-content">{{ info.ja_name ? info.ja_name : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动名称-繁体中文</div>
                                <div class="level-detail-div-item-content">{{ info.zh_CN_name ? info.zh_CN_name : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动名称-English</div>
                                <div class="level-detail-div-item-content">{{ info.en_name ? info.en_name : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动名称-한국어</div>
                                <div class="level-detail-div-item-content">{{ info.ko_name ? info.ko_name : '--' }}</div>
                                </div>
                            </n-gi>
                        </n-grid>
                    </div>
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            基本信息
                        </div>
                        <n-grid y-gap="25" :cols="2">
                            <n-gi>
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">活动封面</div>
                                    <div class="level-detail-div-item-content" v-if="info.cover">
                                        <n-image style="width: 100px" :src="info.cover"/>
                                    </div>
                                    <div class="level-detail-div-item-content" v-else>
                                        --
                                    </div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">活动有效期</div>
                                    <div class="level-detail-div-item-content">{{ info.validityType == 1 ? info.startTime.split(' ')[0] + ' ~ ' + info.endTime.split(' ')[0] : '长期有效' }}</div>
                                </div>
                            </n-gi>
                            <n-gi :span="2">
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">活动规则</div>
                                    <div class="level-detail-div-item-content">{{ info.rule ? info.rule : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">是否启用</div>
                                    <div class="level-detail-div-item-content">
                                        <div class="level-detail-div-item-content-tag1 success" v-if="info.isEnabled == 1">正常</div>
                                        <div class="level-detail-div-item-content-tag1 warning" v-else>停用</div>
                                    </div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">状态</div>
                                    <div class="level-detail-div-item-content">
                                         <div class="level-detail-div-item-content-tag1 warning" v-if="info.status == 1">未开始</div>
                                        <div class="level-detail-div-item-content-tag1 success" v-else-if="info.status == 2">进行中</div>
                                        <div class="level-detail-div-item-content-tag1 warning" v-else>已结束</div>
                                    </div>
                                </div>
                            </n-gi>
                        </n-grid>
                    </div>
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            活动描述
                        </div>
                        <n-grid y-gap="25" :cols="1">
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动描述-中文</div>
                                <div class="level-detail-div-item-content">{{ info.zh_desc ? info.zh_desc : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动描述-日本语</div>
                                <div class="level-detail-div-item-content">{{ info.ja_desc ? info.ja_desc : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动描述-繁体中文</div>
                                <div class="level-detail-div-item-content">{{ info.zh_CN_desc ? info.zh_CN_desc : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动描述-English</div>
                                <div class="level-detail-div-item-content">{{ info.en_desc ? info.en_desc : '--' }}</div>
                                </div>
                            </n-gi>
                            <n-gi>
                                <div class="level-detail-div-item">
                                <div class="level-detail-div-item-title">活动描述-한국어</div>
                                <div class="level-detail-div-item-content">{{ info.ko_desc ? info.ko_desc : '--' }}</div>
                                </div>
                            </n-gi>
                        </n-grid>
                    </div>
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            礼品券
                        </div>
                        <n-grid y-gap="25" :cols="1">
                            <n-gi>
                                <n-table>
                                    <thead>
                                        <tr>
                                            <th>礼品券名称</th>
                                            <th>适用商户</th>
                                            <th>每人每天领取数</th>
                                            <th>每人领取数量</th>
                                            <th>每人每天核销数</th>
                                        </tr>
                                    </thead>
                                    <tbody>
                                        <tr v-for="(item, index) of thCouponArr" :key="index">
                                            <td>{{ item.couponName }}</td>
                                            <td>
                                            <n-ellipsis expand-trigger="click" :line-clamp="1" :tooltip="false" style="width: 200px">
                                                <div v-for="(item1, index1) of item.mchList" :key="index1">{{ item1.mchInfo.name }}</div>
                                            </n-ellipsis>
                                            </td>
                                            <td>
                                              {{ item.perNum }}
                                            </td>
                                            <td>
                                                {{ item.num }}
                                            </td>
                                            <td>
                                              {{ item.perVerifyNum }}
                                            </td>
                                        </tr>
                                    </tbody>
                                </n-table>
                            </n-gi>
                          <n-gi>
                            <div class="level-detail-div-item">
                              <div class="level-detail-div-item-title">券有效期</div>
                              <div class="level-detail-div-item-content">{{ info.couponValidity == 1 ? '跟随活动' : '券自身有效期' }}</div>
                            </div>
                          </n-gi>
                        </n-grid>
                    </div>
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            限制类型
                        </div>
                        <n-grid y-gap="25" :cols="1">
                            <n-gi v-if="info.restrictionType == 1">
                                <div>不限制</div>
                            </n-gi>
                            <n-gi v-if="info.restrictionType == 2">
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">限制部门</div>
                                    <div class="level-detail-div-item-content">{{ departmentName }}</div>
                                </div>
                            </n-gi>
                            <n-gi v-if="info.restrictionType == 3">
                                <div class="level-detail-div-item">
                                    <div class="level-detail-div-item-title">限制员工</div>
                                </div>
                                <n-table>
                                <thead>
                                    <tr>
                                    <th>名称</th>
                                    <th>电话</th>
                                    <th>所属部门</th>
                                </tr>
                                </thead>
                                <tbody>
                                    <tr v-for="(item, index) of employeeArr" :key="index">
                                        <td>{{ item.name }}</td>
                                        <td>{{ item.phoneArea + ' ' + item.phone }}</td>
                                        <td>{{ item.departmentDetail.name }}</td>
                                    </tr>
                                </tbody>
                            </n-table>
                            </n-gi>
                        </n-grid>
                    </div>
                    <div class="level-detail-div" style="padding: 0 20px;">
                        <div class="level-detail-div-title">
                            <div></div>
                            领取记录
                        </div>
                        <n-grid y-gap="25" :cols="1">
                            <n-gi>
                                <div>
                                    <div class="order-table-header">
                                        <div style="padding: 0 1.2%;width: 16%">券号</div>
                                        <div style="padding: 0 1.2%;width: 16%">礼品券</div>
                                        <div style="padding: 0 1.2%;width: 16%">领取员工</div>
                                        <div style="padding: 0 1.2%;width: 16%">领取时间</div>
                                        <div style="padding: 0 1.2%;width: 8%">核销状态</div>
                                        <div style="padding: 0 1.2%;width: 29%">核销信息</div>
                                    </div>
                                     <div class="order-table-body">
                                        <div class="order-table-body-item" v-for="(item, index) in thMemberCouponArr" :key="index">
                                            <div style="padding: 0 1.2%;width: 16%">
                                                <n-tooltip trigger="hover" :arrow-point-to-center="true">
                                                    <template #trigger>
                                                    <div style="width: 100%;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{ item.couponNo }}</div>
                                                    </template>
                                                    {{ item.couponNo }}
                                                </n-tooltip>
                                            </div>
                                            <div style="padding: 0 1.2%;width: 16%">
                                                <n-tooltip trigger="hover" :arrow-point-to-center="true">
                                                    <template #trigger>
                                                    <div style="width: 100%;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">{{ item.thCoupon.couponName }}</div>
                                                    </template>
                                                    {{ item.thCoupon.couponName }}
                                                </n-tooltip>
                                            </div>
                                            <div style="padding: 0 1.2%;width: 16%">
                                                <div>{{ item.employeeInfo.name }}</div>
                                                <div>{{ item.employeeInfo.phoneArea + item.employeeInfo.phone }}</div>
                                            </div>
                                            <div style="padding: 0 1.2%;width: 16%">{{ item.createAt }}</div>
                                            <div style="padding: 0 1.2%;width: 8%">
                                                <template v-if="item.state == 1">
                                                    未生效
                                                </template>
                                                <template v-if="item.state == 2">
                                                    未使用
                                                </template>
                                                <template v-if="item.state == 3">
                                                    已核销
                                                </template>
                                                <template v-if="item.state == 4">
                                                    已过期
                                                </template>
                                                <template v-if="item.state == 5">
                                                    已失效
                                                </template>
                                            </div>
                                            <div style="padding: 0 1.2%;width: 29%">
                                                <template v-if="item.state == 3">
                                                    <n-tooltip trigger="hover" :arrow-point-to-center="true">
                                                        <template #trigger>
                                                        <div style="width: 100%;white-space: nowrap;overflow: hidden;text-overflow: ellipsis;">商户：{{ item.verifyMch.name }}...</div>
                                                        </template>
                                                        <div>商户：{{ item.verifyMch.name }}</div>
                                                        <div>门店：{{ item.verifyStore.storeName }}</div>
                                                        <div>核销商品名：{{ item.thCouponMchName }}</div>
                                                        <div>核销时间：{{ item.verifyTime }}</div>
                                                    </n-tooltip>
                                                </template>
                                                <template v-else>
                                                    --
                                                </template>
                                            </div>
                                        </div>
                                     </div>
                                </div>
                            </n-gi>
                        </n-grid>
                    </div>
                </n-spin>
            </n-drawer-content>
        </n-drawer>
    </div>
</template>

<script lang="ts" setup>
import {computed, ref} from "vue";
import {adaModalWidth} from "@/utils/hotgo";
import {jsontoobj} from "@/utils/smjcomm";
import {View} from "@/api/employeeActivity";
import { EmployeeInter } from '/#/employee';
import { ThCouponMchInter } from '/#/thCoupon';
import {All} from "@/api/thCoupon";
import {List as EmployeeList} from "@/api/employee";
import {List as ThMemberCouponList} from "@/api/thMemberCoupon";

interface SelectThCouponInter {
  id: number;
  couponName: string;
  mchList: ThCouponMchInter[];
  num: number;
  perNum: number;
  perVerifyNum: number;
}

interface ThCouponIdsNumInter {
  couponId: number;
  num: number;
  perNum: number;
  perVerifyNum: number;
}

const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});

const activityId = ref(0);
const info = ref({});
const thCouponIdsNumArr = ref<ThCouponIdsNumInter[]>([]);
const thCouponArr = ref<SelectThCouponInter[]>([])
const departmentName = ref('')
const employeeIdsArr = ref<number[]>([]);
const employeeArr = ref<EmployeeInter[]>([])
const thMemberCouponArr = ref([])

async function getInfo (id) {
    const res = await View({ id: id, isLanguage: true });
    res.nameLanguage = jsontoobj(res.nameLanguage);
    res.descriptionLanguage = jsontoobj(res.descriptionLanguage);

    info.value = res;

    info.value.zh_name = res.nameLanguage.zh && res.nameLanguage.zh.content ? res.nameLanguage.zh.content : ''
    info.value.en_name = res.nameLanguage.en && res.nameLanguage.en.content ? res.nameLanguage.en.content : ''
    info.value.ko_name = res.nameLanguage.ko && res.nameLanguage.ko.content ? res.nameLanguage.ko.content : ''
    info.value.ja_name = res.nameLanguage.ja && res.nameLanguage.ja.content ? res.nameLanguage.ja.content : ''
    if(res.nameLanguage.zh_CN){
    info.value.zh_CN_name = res.nameLanguage.zh_CN.content ? res.nameLanguage.zh_CN.content : ''
    }else if(res.nameLanguage.zh_cn){
    info.value.zh_CN_name = res.nameLanguage.zh_cn.content ? res.nameLanguage.zh_cn.content : ''
    }else{
    info.value.zh_CN_name = ''
    }

    info.value.zh_desc = res.descriptionLanguage.zh && res.descriptionLanguage.zh.content ? res.descriptionLanguage.zh.content : ''
    info.value.en_desc = res.descriptionLanguage.en && res.descriptionLanguage.en.content ? res.descriptionLanguage.en.content : ''
    info.value.ko_desc = res.descriptionLanguage.ko && res.descriptionLanguage.ko.content ? res.descriptionLanguage.ko.content : ''
    info.value.ja_desc = res.descriptionLanguage.ja && res.descriptionLanguage.ja.content ? res.descriptionLanguage.ja.content : ''
    if(res.descriptionLanguage.zh_CN){
    info.value.zh_CN_desc = res.descriptionLanguage.zh_CN.content ? res.descriptionLanguage.zh_CN.content : ''
    }else if(res.descriptionLanguage.zh_cn){
    info.value.zh_CN_desc = res.descriptionLanguage.zh_cn.content ? res.descriptionLanguage.zh_cn.content : ''
    }else{
    info.value.zh_CN_desc = ''
    }

    thCouponArr.value = []
    thCouponIdsNumArr.value = []
    if(res.couponList){
        res.couponList.forEach((item) => {
            thCouponIdsNumArr.value.push({couponId: item.couponId, num: item.availableQuantity, perNum: item.perDayAvailable,  perVerifyNum: item.perDayVerify})
        })
    }

    if(res.restrictionType == 2){
      departmentName.value = res.employeeDepartmentList.map(m => m.departmentDetail.name).join(',')
    }

    if(res.restrictionType == 3){
        employeeIdsArr.value = res.employeeList.map(m => m.employeeId);
        employeeArr.value = [];
        await getEmployeeList()
    }
};

async function getThCouponList(){
  const res = await All({
    couponIds: thCouponIdsNumArr.value.map(m => m.couponId).join(',')
  });
  thCouponArr.value = res.list
  if(thCouponArr.value.length > 0){
      thCouponArr.value.forEach(m => {
        m.num = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.num || 1;
        m.perNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perNum || 0
        m.perVerifyNum = thCouponIdsNumArr.value.find(n => n.couponId == m.id)?.perVerifyNum || 0
      })
    }
  console.log('thCouponArr',thCouponArr.value)
}

async function getEmployeeList(){
  const res = await EmployeeList({
    employeeIds: employeeIdsArr.value.join(','),
    pagination: false
  });
  employeeArr.value = res.list
  console.log('employeeArr',employeeArr.value)
}

async function getThMemberCouponList(){
  const res = await ThMemberCouponList({
    activityId: activityId.value,
    pagination: false
  });
  thMemberCouponArr.value = res.list
  console.log('thMemberCouponArr',thMemberCouponArr.value)
}

async function openModal(id) {
  activityId.value = id
  showModal.value = true;
  loading.value = true;
  await getInfo(id);
  await getThCouponList()
  await getThMemberCouponList()
  loading.value = false;
}

defineExpose({
  openModal,
});
</script>

<style lang="less" scoped>
::v-deep(.json-width) {
  width: 100%;
  min-width: 3.125rem;
}
.level-detail-div{
  margin-bottom: 25px;
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 25px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 12px;
      color: #707070;
      line-height: 17px;
      margin-bottom: 8px;
    }
    &-content{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      display: flex;
      align-items: center;
      &-tag{
        width: 38px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;
        &.success{
          color: #17A158;
          border: 1px solid #17A158;
        }
        &.warning{
          color: #F48720;
          border: 1px solid #F48720;
        }
      }
      &-tag1{
        padding: 0 5px;
        height: 22px;
        line-height: 22px;
        text-align: center;
        font-size: 14px;
        border-radius: 2px;
        font-weight: 400;
        &.success{
          color: #26A763;
          background: #E3F4EB;
        }
        &.warning{
          color: #EFA020;
          background: #FFECCE;
        }
        &.primary{
          color: #3F9EFF;
          background: #ECF5FF;
        }
        &.error{
          color: #F56C6C;
          background: #FEF0F0;
        }
      }
    }
  }
}
.coupon-statist-div{
  display: flex;
  align-items: center;
  justify-content: space-between;
  &-item{
    width: 340px;
    height: 88px;
    background: #F9FBFC;
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 0 15px;
    div{
      &:first-child{
        font-weight: 400;
        font-size: 14px;
        color: #707070;
        line-height: 20px;
      }
      &:last-child{
        margin-top: 5px;
        font-weight: 500;
        font-size: 24px;
        color: #3D3D3D;
        line-height: 33px;
      }
    }
  }
}
.car-order-detail-tabs{
  background: #F9FBFC;
  height: 40px;
  line-height: 40px;
  padding-left: 20px;
  display: flex;
  &-item{
    height: 40px;
    line-height: 40px;
    position: relative;
    font-weight: 500;
    font-size: 14px;
    color: #3D3D3D;
    margin-right: 36px;
    cursor: pointer;
    &.active{
      color: #053DC8;
      div{
        position: absolute;
        left: 0;
        right: 0;
        bottom: -1px;
        background: #053DC8;
        height: 2px;
      }
    }
  }
}

.order-table-header{
    display: flex;
    align-items: center;
    padding: 10px 0;
    line-height: 40px;
    background: #F7F8FA;
    div{
      font-weight: 500;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 40px;
    }
  }
  .order-table-body{
    &-item{
      display: flex;
      align-items: center;
      border-bottom: 1px solid #EEEEEE;
      padding: 10px 0;
      &:hover{
        background: #FAFAFC;
      }
      div{
        font-weight: 400;
        font-size: 14px;
        color: #3D3D3D;
      }
      background: #FFFFFF;
      &:nth-child(even) {
        background: #FAFAFC; /* 例如，将背景色设置为灰色 */
      }
    }
  }
</style>
