const { rule } = require("postcss");

module.exports = {
	root: true,
	extends: ['eslint:recommended', 'prettier'],
	plugins: ['svelte3'],
	overrides: [{ files: ['*.svelte'], processor: 'svelte3/svelte3' }],
	parserOptions: {
		sourceType: 'module',
		ecmaVersion: 2020
	},
	env: {
		browser: true,
		es2017: true,
		node: true
	},
	rules: {
		"prettier/prettier": [
			"error",
			{
				"endOfLine": "auto",
				"semi": false

			}
		],
		"no-unused-vars": "warn",
		"linebreak-style": ["error", "windows"],
		"no-console": "off",
		"quotes": ["single", { "allowTemplateLiterals": true }],
		"indent": ["error", 4],
		"semi": ["error", "never"]
	}
}
