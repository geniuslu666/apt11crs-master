<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card
        :bordered="false"
        class="proCard"
        :header-style="{
                    padding: '25px 20px 20px',
                  }"
        :content-style="{
                    padding: '0 20px 20px',
                  }"
      >
        <template #header>
          <text style="font-weight: 500;font-size: 20px;color: #3D3D3D;line-height: 28px;">配置信息</text>
        </template>
        <div class="statusTab">
          <div :class="tab == 0 ? 'active' : ''"><span @click="handleTab(0)">基本设置</span></div>
          <div :class="tab == 1 ? 'active' : ''"><span @click="handleTab(1)">微信分享配置</span></div>
          <div :class="tab == 2 ? 'active' : ''"><span @click="handleTab(2)">支付方式</span></div>
          <div :class="tab == 3 ? 'active' : ''"><span @click="handleTab(3)">民宿价格配置</span></div>
          <div :class="tab == 4 ? 'active' : ''"><span @click="handleTab(4)">首页导航配置</span></div>
          <div :class="tab == 5 ? 'active' : ''"><span @click="handleTab(5)">汇率配置</span></div>
          <div :class="tab == 6 ? 'active' : ''"><span @click="handleTab(6)">语言包配置</span></div>
          <div :class="tab == 7 ? 'active' : ''"><span @click="handleTab(7)">测试专区配置</span></div>
          <div :class="tab == 9 ? 'active' : ''"><span @click="handleTab(9)">储物柜Api</span></div>
          <div :class="tab == 10 ? 'active' : ''"><span @click="handleTab(10)">ToretaApi</span></div>
          <div :class="tab == 8 ? 'active' : ''"><span @click="handleTab(8)">小程序审核</span></div>
          <div :class="tab == 11 ? 'active' : ''"><span @click="handleTab(11)">联系信息</span></div>
          <div :class="tab == 12 ? 'active' : ''"><span @click="handleTab(12)">首页广告位</span></div>
          <div :class="tab == 13 ? 'active' : ''"><span @click="handleTab(13)">个人中心按钮</span></div>
        </div>
        <AppSetting ref="appSettingRef" v-if="tab == 0"/>
        <WxShareSetting ref="wxShareSettingRef" v-if="tab == 1"/>
        <PaymodeConfig ref="paymodeConfigRef" v-if="tab == 2"/>
        <PmsPrice ref="pmsPriceRef" v-if="tab == 3"/>
        <IndexNav ref="indexNav" v-if="tab == 4"/>
        <ExchangeRate ref="exchangeRate" v-if="tab == 5"/>
        <LanguagePack ref="languagePack" v-if="tab == 6"/>
        <TestNav ref="testNav" v-if="tab == 7"/>
        <WxMiniAudit ref="wxMiniAudit" v-if="tab == 8"/>
        <CabinetApi ref="cabinetApi" v-if="tab == 9"/>
        <ToretaApi ref="toretaApi" v-if="tab == 10"/>
        <ContactSetting ref="contactSetting" v-if="tab == 11"/>
        <IndexBanner ref="indexBanner" v-if="tab == 12"/>
        <ProfileButtonSetting ref="profileButtonSetting" v-if="tab == 13"/>
      </n-card>
    </n-spin>
  </div>
</template>
<script lang="ts" setup>
import {ref} from 'vue';
import AppSetting from "@/views/appsystem/appSetting.vue";
import WxShareSetting from "@/views/appsystem/wxShareSetting.vue";
import PaymodeConfig from "@/views/appsystem/paySetting.vue";
import PmsPrice from "@/views/appsystem/pmsPrice.vue";
import IndexNav from "@/views/appsystem/indexNav.vue";
import ExchangeRate from "@/views/appsystem/exchangeRate.vue";
import LanguagePack from "@/views/appsystem/languagePack.vue";
import TestNav from "@/views/appsystem/testNav.vue";
import WxMiniAudit from "@/views/appsystem/wxMiniAudit.vue";
import CabinetApi from "@/views/appsystem/cabinetApi.vue";
import ToretaApi from "@/views/appsystem/toretaApi.vue";
import ContactSetting from "@/views/appsystem/contactSetting.vue";
import IndexBanner from "@/views/appsystem/indexBanner.vue";
import ProfileButtonSetting from "@/views/appsystem/profileButtonSetting.vue";

const show = ref(false);
const tab = ref(0);
const appSettingRef = ref();
const wxShareSettingRef = ref();
const paymodeConfigRef = ref();
const pmsPriceRef = ref();
const testNav = ref();
const wxMiniAudit = ref();
const cabinetApi = ref();
const toretaApi = ref();
const contactSetting = ref();

function handleTab(e){
  tab.value = e;
}
</script>
<style lang="less" scoped>
.statusTab{
  display: flex;
  margin-bottom: 4px;
  div{
    margin-right: 5px;
    padding: 0 16px;
    height: 30px;
    line-height: 30px;
    text-align: center;
    font-size: 14px;
    color: #4E5969;
    span{
      cursor: pointer;
    }
    &.active{
      background: #F2F3F8;
      border-radius: 30px;
      color: #1664FF;
      font-weight: 500;
    }
  }
}
</style>
