<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '25px 0 0',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">门店详情</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            返回
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="editBtn" style="width: 70px;height: 35px;">
            编辑
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              门店概要
            </div>
            <div class="coupon-statist-div">
              <div class="coupon-statist-div-item">
                <div style="display: flex; justify-content: space-between; align-items: center; width: 100%;">
                  <div style="display: inline-block;">已核销</div>
                  <n-button type="primary" ghost @click="handleExport" class="min-left-space" size="small" strong
                            v-if="hasPermission(['/thMemberCoupon/export'])">
                    导出
                  </n-button>
                </div>
                <div>
                  {{ info.verifyNum }}

                </div>
              </div>
            </div>
          </div>
          <div class="level-detail-div" style="padding: 0 20px;">
            <div class="level-detail-div-title">
              <div></div>
              门店信息
            </div>
            <n-grid y-gap="25" :cols="3">
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">所属商户</div>
                  <div class="level-detail-div-item-content">{{ info.thMchName }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店名称-中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_name ? info.zh_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店名称-日本语</div>
                  <div class="level-detail-div-item-content">{{ info.ja_name ? info.ja_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店名称-繁体中文</div>
                  <div class="level-detail-div-item-content">{{ info.zh_CN_name ? info.zh_CN_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店名称-English</div>
                  <div class="level-detail-div-item-content">{{ info.en_name ? info.en_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi>
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店名称-한국어</div>
                  <div class="level-detail-div-item-content">{{ info.ko_name ? info.ko_name : '--' }}</div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店图片</div>
                  <div class="level-detail-div-item-content">
                    <n-image-group>
                      <n-space>
                        <span v-for="(item, key) in info?.imagesArr" :key="key">
                          <n-image style="margin-left: 10px; height: 100px; width: 100px" :src="item" />
                        </span>
                      </n-space>
                    </n-image-group>
                  </div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店电话</div>
                  <div class="level-detail-div-item-content">{{ info.phone ? info.phone : '--' }}</div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">门店地址</div>
                  <div class="level-detail-div-item-content">{{ info.detailAddress }}</div>
                </div>
                <div class="level-detail-div-item" style="position: relative">
                  <!-- 地址详细信息 -->
                  <div class="mapbtm f12" v-if="info.detailAddress">
                    {{ info.detailAddress }}
                  </div>
                  <div class="mapbtm f12" v-else> {{ translang('暂无定位信息') }}</div>

                  <img
                    style="width: 100%; height: 260px; object-fit: cover"
                    class="mt-3 mb-5"
                    src="@/assets/images/nomap.png"
                    v-if="!info.ggLat || info.ggLat == 'undefined' || showEdit"
                  />
                  <GMapView
                    class="mt-3"
                    :lat="info.ggLat"
                    :long="info.ggLng"
                    :heightmap="260"
                    v-else
                  />
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">账号</div>
                  <div class="level-detail-div-item-content">{{ info.account }}</div>
                </div>
              </n-gi>
              <n-gi span="3">
                <div class="level-detail-div-item">
                  <div class="level-detail-div-item-title">绑定终端</div>
                </div>
                <n-table>
                  <thead>
                  <tr>
                    <th>终端编号</th>
                    <th>终端名称</th>
                    <th>终端类型</th>
                    <th>品牌型号</th>
                    <th>打印联数</th>
                  </tr>
                  </thead>
                  <tbody>
                  <tr v-for="(item, index) of terminalArr" :key="index">
                    <td>{{ item.terminalInfo.sn }}</td>
                    <td>{{ item.terminalInfo.terminalName }}</td>
                    <td>
                      <n-tag
                        :type="getOptionTag(options.terminal_type, item.terminalInfo.terminalType)"
                        size="small"
                        class="min-left-space"
                      >
                        {{ getOptionLabel(options.terminal_type, item.terminalInfo.terminalType) }}
                      </n-tag>
                    </td>
                    <td>
                      <template v-if="item.terminalInfo.terminalType === 'VERIFY_PRINTER'">
                        <n-tag
                          :type="getOptionTag(options.verify_brand_model, item.terminalInfo.brandModel)"
                          size="small"
                          class="min-left-space"
                        >
                          {{ getOptionLabel(options.verify_brand_model, item.terminalInfo.brandModel) }}
                        </n-tag>
                      </template>
                      <template v-else-if="item.terminalInfo.terminalType === 'HAND_TERMINAL'">
                        <n-tag
                          :type="getOptionTag(options.hand_brand_model, item.terminalInfo.brandModel)"
                          size="small"
                          class="min-left-space"
                        >
                          {{ getOptionLabel(options.hand_brand_model, item.terminalInfo.brandModel) }}
                        </n-tag>
                      </template>
                    </td>
                    <td>{{ item.terminalInfo.terminalType == 'VERIFY_PRINTER' ? item.printTimes : '--' }}</td>
                  </tr>
                  </tbody>
                </n-table>
              </n-gi>
            </n-grid>
          </div>
        </n-spin>
      </n-drawer-content>
    </n-drawer>
    <Edit ref="editRef" @reloadTable="reloadTable" @closeEdit="handleCloseEdit"/>
  </div>
</template>

<script setup lang="ts">
import {computed, ref} from "vue";
import {adaModalWidth, getOptionLabel, getOptionTag} from "@/utils/hotgo";
import {View} from "@/api/thMchStore";
import {jsontoobj, translang} from "@/utils/smjcomm";
import GMapView from "@/views/smjcomm/GMapView.vue";
import {options} from "@/views/terminal/model";
import Edit from "@/views/thMchStore/edit.vue";
import {ExportOutlined} from "@vicons/antd";
import {NButton, useMessage} from "naive-ui";
import {usePermission} from "@/hooks/web/usePermission";
import { Export } from '@/api/thMemberCoupon';

const emit = defineEmits(['reloadTable']);
const formBtnLoading = ref(false);
const showModal = ref(false);
const loading = ref(false);
const showEdit = ref(false)
const { hasPermission } = usePermission();
const message = useMessage();
const dialogWidth = computed(() => {
  return adaModalWidth(1100);
});

const editRef = ref();
const storeId = ref(0);
const info = ref({});
const terminalArr = ref([])

async function getInfo(id){
  const res = await View({
    id: id,
    isLanguage: true
  });
  res.nameLanguage = jsontoobj(res.nameLanguage);

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

  info.value.imagesArr = res.images ? res.images.split(',').map((item)=>{
    return item
  }) : [];

  terminalArr.value = []

  if(res.terminalList && res.terminalList.length > 0){
    res.terminalList.forEach((item,index) => {
      if(item.terminalInfo){
        terminalArr.value.push(item)
      }
    })
  }

}

async function reloadTable(){
  formBtnLoading.value = true
  loading.value = true;
  await getInfo(storeId.value);
  loading.value = false;
  formBtnLoading.value = false
  emit('reloadTable');
}

async function openModal(id) {
  storeId.value = id
  showModal.value = true;
  formBtnLoading.value = true
  loading.value = true;
  await getInfo(id);
  loading.value = false;
  formBtnLoading.value = false
}

function closeForm() {
  showModal.value = false;
  loading.value = false;
  formBtnLoading.value = false
}

function editBtn() {
  showEdit.value = true
  editRef.value.openModal(info.value);
}

function handleCloseEdit(){
  showEdit.value = false
}

// 导出
function handleExport() {
  message.loading('正在导出列表...', { duration: 2000 });
  Export({ verifyStoreId: storeId.value, state: 3 });
}


defineExpose({
  openModal,
});
</script>

<style scoped lang="less">
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
.mapbtm {
  position: absolute;
  z-index: 300;
  left: 0;
  right: 0;
  bottom: 0;
  text-align: left;
  background: #6a7b8c38;
  padding: 5px 10px;
  font-weight: 550;
}
</style>
