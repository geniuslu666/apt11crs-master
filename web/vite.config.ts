import type {ConfigEnv, UserConfig} from 'vite';
import {loadEnv} from 'vite';
import {resolve} from 'path';
import {wrapperEnv} from './build/utils';
import {createVitePlugins} from './build/vite/plugin';
import {OUTPUT_DIR} from './build/constant';
import {createProxy} from './build/vite/proxy';
import pkg from './package.json';
import {format} from 'date-fns';

const { dependencies, devDependencies, name, version } = pkg;

const __APP_INFO__ = {
  pkg: { dependencies, devDependencies, name, version },
  lastBuildTime: format(new Date(), 'yyyy-MM-dd HH:mm:ss'),
};

function pathResolve(dir: string) {
  return resolve(process.cwd(), '.', dir);
}

export default ({ command, mode }: ConfigEnv): UserConfig => {
  const root = process.cwd();
  const env = loadEnv(mode, root);
  const viteEnv = wrapperEnv(env);
  const { VITE_PUBLIC_PATH, VITE_PORT, VITE_PROXY } = viteEnv;
  const isBuild = command === 'build';
  return {
    base: VITE_PUBLIC_PATH,
    esbuild: {
      drop: ['debugger'],
      pure: ['console.log'],
      // 禁用代码检查
      legalComments: 'none'
    },
    resolve: {
      alias: [
        {
          find: /\/#\//,
          replacement: pathResolve('types') + '/',
        },
        {
          find: '@',
          replacement: pathResolve('src') + '/',
        },
      ],
      dedupe: ['vue'],
    },
    plugins: createVitePlugins(viteEnv, isBuild),
    define: {
      __APP_INFO__: JSON.stringify(__APP_INFO__),
      __VUE_OPTIONS_API__: true,
      __VUE_PROD_DEVTOOLS__: false,
      __VUE_PROD_HYDRATION_MISMATCH_DETAILS__: true,
    },
    css: {
      preprocessorOptions: {
        less: {
          modifyVars: {},
          javascriptEnabled: true,
          additionalData: `@import "src/styles/var.less";`,
        },
      },
    },
    server: {
      host: true,
      port: VITE_PORT,
      proxy: createProxy(VITE_PROXY),
      // proxy: {
      //     '/api': {
      //         target: '',
      //         changeOrigin: true,
      //         rewrite: (path) => path.replace(/^\/api/, '/api/v1')
      //     }
      // }
    },
    optimizeDeps: {
      include: [],
      exclude: ['vue-demi'],
    },
    build: {
      target: 'esnext',
      cssTarget: 'chrome80',
      outDir: OUTPUT_DIR,
      // 关闭 gzip 大小报告，可节省大量构建时间
      reportCompressedSize: false,
      chunkSizeWarningLimit: 3000,
      // 启用 esbuild 压缩（比 terser 快 20-40 倍）
      minify: 'esbuild',
      // 禁用 sourcemap 加速构建
      sourcemap: false,
      rollupOptions: {
        // 降低内存占用：禁用树摇优化的缓存
        cache: false,
        onwarn(warning, warn) {
          // 忽略 eslint 相关警告
          if (warning.plugin === 'eslint') return;
          warn(warning);
        },
        output: {
          // 手动分包，减少单个 chunk 大小，降低内存峰值
          manualChunks: {
            'vue-vendor': ['vue', 'vue-router', 'pinia'],
            'naive-ui': ['naive-ui'],
            'echarts': ['echarts'],
            'ant-design': ['ant-design-vue'],
            'tinymce': ['tinymce', '@tinymce/tinymce-vue'],
          },
        },
      },
    },
  };
};
