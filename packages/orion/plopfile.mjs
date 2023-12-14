import { username } from 'username'
import dayjs from 'dayjs'

const defaultValue = 'unknown'

export default function (
  /** @type {import('plop').NodePlopAPI} */
  plop
) {
  plop.setGenerator('adr', {
    description: 'docs/adrにADRを記述するためのファイルを生成します。',
    prompts: [
      {
        type: 'input',
        name: 'title',
        message: 'タイトルを入力してください。',
      },
      {
        type: 'list',
        name: 'tag',
        message: 'タグを選択してください。',
        choices: [
          {
            name: 'general',
            value: {
              name: 'general',
              omitted: 'ge'
            },
          },
          {
            name: 'front-end',
            value: {
              name: 'front-end',
              omitted: 'fe'
            },
          },
          {
            name: 'back-end',
            value: {
              name: 'back-end',
              omitted: 'be'
            },
          },
        ]
      }
    ],
    actions: [
      {
        type: 'add',
        path: 'docs/adr/{{dateFilename}}_{{tag.omitted}}_{{title}}.md',
        templateFile: 'plop-templates/adr.md.hbs',
        data: async () => {
          const now = dayjs()
          const name = await username()


          return {
            date: now.format('YYYY/MM/DD HH:mm:ss'),
            dateFilename: now.format('YYYYMMDDHHmm'),
            username: name ?? defaultValue,
          }
        }
      }
    ]
  })
}