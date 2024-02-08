module.exports = {
  theme: 'cosmos',
  title: 'Tenderhellar',
  // locales: {
  //   "/": {
  //     lang: "en-US"
  //   },
  //   "/ru/": {
  //     lang: "ru"
  //   }
  // },
  base: process.env.VUEPRESS_BASE,
  themeConfig: {
    repo: 'hellarpro/tenderhellar',
    docsRepo: 'hellarpro/tenderhellar',
    docsDir: "docs",
    editLinks: true,
    label: 'core',
    algolia: {
      id: "BH4D9OD16A",
      key: "59f0e2deb984aa9cdf2b3a5fd24ac501",
      index: "tenderhellar"
    },
    versions: [
      {
        "label": "v0.7",
        "key": "v0.7"
      },
      {
        "label": "v0.8",
        "key": "v0.8"
      },
      {
        "label": "v0.9",
        "key": "v0.9"
      }
    ],
    topbar: {
      banner: false,
    },
    sidebar: {
      auto: true,
      nav: [
        {
          title: 'Resources',
          children: [
            {
              title: 'Developer Sessions',
              path: '/DEV_SESSIONS.html'
            },
            {
              // TODO(creachadair): Figure out how to make this per-branch.
              // See: https://github.com/tendermint/tendermint/issues/7908
              title: 'RPC',
              path: 'https://docs.tendermint.com/v0.35/rpc/',
              static: true
            },
          ]
        }
      ]
    },
    gutter: {
      title: 'Help & Support',
      editLink: true,
      forum: {
        title: 'Tenderhellar Forum',
        text: 'Join the Tenderhellar forum to learn more',
        url: 'https://forum.cosmos.network/c/tendermint',
        bg: '#0B7E0B',
        logo: 'tenderhellar'
      },
      github: {
        title: 'Found an Issue?',
        text: 'Help us improve this page by suggesting edits on GitHub.'
      }
    },
    footer: {
      question: {
        text: 'Chat with Tenderhellar developers in <a href=\'https://discord.gg/fqfCb4fX\' target=\'_blank\'>Discord</a>'
      },
      logo: '/logo-bw.svg',
      textLink: {
        text: 'hellar.io',
        url: 'https://hellar.io'
      },
      services: [
        {
          service: 'medium',
          url: 'https://medium.com/@hellarpay'
        },
        {
          service: 'twitter',
          url: 'https://twitter.com/hellarpay'
        },
        {
          service: 'linkedin',
          url: 'https://www.linkedin.com/company/hellar-core-group/'
        },
        {
          service: 'reddit',
          url: 'https://reddit.com/r/hellarpay'
        },
        {
          service: 'telegram',
          url: 'https://t.me/hellar_chat'
        },
        {
          service: 'youtube',
          url: 'https://www.youtube.com/channel/UCAzD2v9Yx4a4iS2_-unODkA'
        }
      ],
      smallprint:
        'The development of Tenderhellar is led by [Hellar Core Group](https://hellar.io/). Funding for this development comes primarily from the Hellar Governance system. The Tendermint trademark is owned by Tendermint Inc. The Tenderhellar logo is owned by Hellar Core Group Inc., the company that maintains this website.',
      links: [
        {
          title: 'Documentation',
          children: [
            {
              title: 'Hellar developer docs',
              url: 'https://hellarplatform.readme.io/'
            },
            {
              title: 'Hellar user docs',
              url: 'https://docs.hellar.io/en/stable/'
            }
          ]
        },
        {
          title: 'Community',
          children: [
            {
              title: 'Hellar blog',
              url: 'https://medium.com/@hellarpay'
            },
            {
              title: 'Forum',
              url: 'https://forum.hellar.io'
            }
          ]
        },
        {
          title: 'Contributing',
          children: [
            {
              title: 'Contributing to the docs',
              url: 'https://github.com/hellarpro/tenderhellar'
            },
            {
              title: 'Source code on GitHub',
              url: 'https://github.com/hellarpro/tenderhellar'
            },
            {
              title: 'Careers at Hellar Core Group',
              url: 'https://hellar.io/dcg/jobs/'
            }
          ]
        }
      ]
    }
  },
  plugins: [
    [
      '@vuepress/google-analytics',
      {
        ga: 'UA-51029217-11'
      }
    ],
    [
      '@vuepress/plugin-html-redirect',
      {
        countdown: 0
      }
    ]
  ]
};
