<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">员工详情</div>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              基本信息
            </div>
            <n-grid y-gap="25" :cols="2">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工姓名</div>
                  <div class="level-detail-div-item-content">
                    {{ data.name }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">员工部门</div>
                  <div class="level-detail-div-item-content">
                    {{ data.departmentDetail ? data.departmentDetail.name : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">手机号</div>
                  <div class="level-detail-div-item-content">
                    {{ data.phoneArea }}-{{ data.phone }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">绑定会员</div>
                  <div class="level-detail-div-item-content">
                    {{ data.memberDetail ? data.memberDetail.fullName : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">状态</div>
                  <div class="level-detail-div-item-content">
                    <div class="level-detail-div-item-content-tag1 success" v-if="data.status == 1">正常</div>
                    <div class="level-detail-div-item-content-tag1 warning" v-else>停用</div>
                  </div>
                </div>
              </n-gi>
              <n-gi span="2">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">备注</div>
                  <div class="level-detail-div-item-content">
                    {{ data.remark ? data.remark : '--' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
            <div class="level-detail-div-title" v-if="data.memberId>0" style="margin-top: 20px">
              <div></div>
              绑定信息
            </div>
            <n-grid y-gap="25" :cols="2" v-if="data.memberId>0">
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员编号</div>
                  <div class="level-detail-div-item-content">
                    {{ data.memberDetail ? data.memberDetail.memberNo : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">会员名称</div>
                  <div class="level-detail-div-item-content">
                    {{ data.memberDetail ? data.memberDetail.fullName : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">手机</div>
                  <div class="level-detail-div-item-content">
                    {{ data.memberDetail ? data.memberDetail.phoneArea + '-' +data.memberDetail.phone : '--' }}
                  </div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">邮箱</div>
                  <div class="level-detail-div-item-content">
                    {{ data.memberDetail ? data.memberDetail.mail : '--' }}
                  </div>
                </div>
              </n-gi>
            </n-grid>
            <div class="level-detail-div-title" style="margin-top: 25px;margin-bottom: 0">
              <div></div>
              礼品券
            </div>
            <BasicTable  ref="actionRef" :columns="columns" :request="loadDataTable" :scroll-x="scrollX" :resizeHeightOffset="-10000">
            </BasicTable>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>

    <EditBase ref="editBaseRef" @reloadInfo="reloadInfo" />
  </div>
</template>

<script lang="ts" setup>
import {computed, onMounted, ref, h} from 'vue';
import { View } from '@/api/employee';
import {BasicTable} from "@/components/Table";
import {adaTableScrollX, adaModalWidth} from "@/utils/hotgo";
import EditBase from "@/views/pmsStaff/edit_base.vue";
import {List} from '@/api/thMemberCoupon';

const emit = defineEmits(['reloadTable']);
const showModal = ref(false);
const loading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(900);
});
const employeeId = ref(0)
// 字典数据选项
const data = ref({});
const editBaseRef = ref();

const getInfo = async () => {
  loading.value = true;
  const res =  await View({id: employeeId.value})
  data.value = res;
  loading.value = false;
};

// 表格列
const columns = [
  {
    title: '礼品券编号',
    key: 'couponNo',
    align: 'left',
    width: 155,
  },
  {
    title: '券名称',
    key: 'couponName',
    align: 'left',
    width: 155,
    render(row){
      return row.thCoupon?.couponName
    }
  },
  {
    title: '核销状态',
    key: 'state',
    align: 'left',
    width: 90,
    ellipsis: false,
    render(row) {
      if(row.state == 1){
        return '未生效'
      }else if(row.state == 2){
        return '未使用'
      }else if(row.state == 3){
        return '已核销'
      }else if(row.state == 4){
        return '已过期'
      }else if(row.state == 5){
        return '已失效'
      }
    }
  },
  {
    title: '核销信息',
    key: 'verifyTime',
    align: 'left',
    width: 165,
    render(row){
      if(!row.verifyTime){
        return '--'
      }
      return h(
        'div',
        {
          style: {
            // display: 'flex',
            // alignItems: 'center'
            lineHeight: "15px"
          }
        },
        [
          h(
            'p',
            {},
            {
              default: () => "核销时间："+row.verifyTime,
            }
          ),
          h(
            'p',
            {},
            {
              default: () => "核销商户："+row.verifyMch?.name,
            }
          ),
          h(
            'p',
            {},
            {
              default: () => "核销门店："+row.verifyStore?.storeName,
            }
          ),
          h(
            'p',
            {},
            {
              default: () => "核销商品名："+row.thCouponMchName,
            }
          ),
        ]
      )
    }
  },
];

// 加载表格数据
const loadDataTable = async (res) => {
  return await List({ ...{
      employeeId: employeeId.value
    }, ...res });
};

const scrollX = computed(() => {
  return adaTableScrollX(columns, 0);
});

async function reloadInfo(){
  await getInfo();
  emit('reloadTable');
}

async function openModal(id) {
  employeeId.value = id
  showModal.value = true;
  await getInfo();
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
      a{
        font-weight: 400;
        font-size: 14px;
        color: #1577FF;
        line-height: 22px;
        margin-left: 6px;
        cursor: pointer;
      }
      .color{
        width: 20px;
        height: 20px;
        margin-right: 8px;
        border-radius: 2px;
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
</style>


