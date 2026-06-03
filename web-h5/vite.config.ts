import { defineConfig, loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import { VantResolver } from '@vant/auto-import-resolver'
import pxToViewport from 'postcss-px-to-viewport-8-plugin'
import tailwindcss from 'tailwindcss'
import autoprefixer from 'autoprefixer'

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  const env = loadEnv(mode, process.cwd())

  // 解析代理配置
  const proxyConfig: Record<string, any> = {}
  if (env.VITE_PROXY) {
    try {
      const proxyList: [string, string][] = JSON.parse(env.VITE_PROXY)
      proxyList.forEach(([prefix, target]) => {
        proxyConfig[prefix] = {
          target,
          changeOrigin: true,
        }
      })
    } catch (e) {
      console.warn('VITE_PROXY 格式错误，请检查 .env 文件')
    }
  }

  return {
    base: '/daytrip/',
    plugins: [
      vue(),
      Components({
        resolvers: [VantResolver()],
      }),
    ],
    css: {
      postcss: {
        plugins: [
          tailwindcss(),
          autoprefixer(),
          pxToViewport({
            viewportWidth: 375,
            unitPrecision: 5,
            viewportUnit: 'vw',
            // Exclude Tailwind classes from px→vw conversion
            selectorBlackList: [/^\.tw-/, /^\.(flex|grid|block|inline|hidden|absolute|relative|fixed|sticky)/],
            minPixelValue: 1,
            mediaQuery: false,
          }),
        ],
      },
    },
    build: {
      outDir: '../resource/public/daytrip',
      emptyOutDir: true,
    },
    server: {
      port: 5177,
      proxy: proxyConfig,
    },
  }
})
