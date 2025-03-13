# tfchain4
**tfchain4** is ThreeFold blockchain for Grid v4. It built using Cosmos SDK and Tendermint.

## Development

### the blockchain

Requirements:
- Go 1.23

**install ignite cli**

```console
curl https://get.ignite.com/cli@v28.0.0! | bash
```

**Start single node chain**

```console
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Vue interface

Requirement:
- nodejs v22

**install dependencies**
```console
cd vue
npm install
```

**run the vuejs interface**

```console
cd vue
npm run dev
```

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).
