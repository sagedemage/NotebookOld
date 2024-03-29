module.exports = {
    "settings": {
        "react": {
          "version": "detect"
        }
    },
    "env": {
        "browser": true,
        "es2021": true
    },
    "extends": [
        "eslint:recommended",
        "plugin:react/recommended",
        "plugin:@typescript-eslint/recommended"
    ],
    "overrides": [
    ],
    "parser": "@typescript-eslint/parser",
    "parserOptions": {
        "ecmaVersion": "latest",
        "sourceType": "module",
        "ecmaFeatures": {
            "jsx": true
        }
    },
    "plugins": [
        "react",
        "@typescript-eslint"
    ],
    "rules": {
        "react/react-in-jsx-scope": "off",
        "react/jsx-uses-react": "off",
        "@typescript-eslint/no-non-null-assertion": "error",
        "eqeqeq": "error",
        "curly": "error",
        "quotes": [
            "error", 
            "double"
        ],
        "no-unused-vars": [
            "error",
            {
                "vars": "all",
                "args": "after-used",     
                "ignoreRestSiblings": false 
             }
          ]
    }
}
