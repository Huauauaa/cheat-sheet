import { defineUserConfig, defaultTheme } from 'vuepress';
import { searchPlugin } from '@vuepress/plugin-search';

export default defineUserConfig({
  base: '/cheat-sheet/',
  lang: 'en-US',
  title: 'Cheat Sheet',
  description: 'This is a cheat sheet',
  theme: defaultTheme({
    navbar: [
      {
        text: 'Home',
        link: '/',
      },
    ],
    sidebar: [
      {
        text: 'Guide',
        link: '/guide/',
        children: [
          '/guide',
          '/book-review',
          '/fiction',
          {
            text: 'issue',
            link: 'https://github.com/Huauauaa/cheat-sheet/issues',
            children: [],
          },
        ],
      },
    ],
  }),
  plugins: [searchPlugin({})],
});
