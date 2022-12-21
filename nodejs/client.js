const PROTO_PATH = './proto/laptop_service.proto'
console.log(PROTO_PATH)

const parseArgs = require('minimist')
const grpc = require('@grpc/grpc-js')
const protoLoader = require('@grpc/proto-loader')
// const { timeout } = require('async');
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
})
const protoDescriptor = grpc.loadPackageDefinition(packageDefinition)
const laptopService = protoDescriptor.lemonade // lemonade is package name in proto file

async function main() {
    const argv = parseArgs(process.argv.slice(2), {
        string: 'addr',
    })
    let target
    if (argv.target) {
        target = argv.target
    } else {
        target = 'localhost:50051'
    }
    //   const client = new hello_proto.Route(target, grpc.credentials.createInsecure());

    //
    const client = new laptopService.LaptopService(
        target,
        grpc.credentials.createInsecure()
    )
    let user
    if (argv._.length > 0) {
        user = argv._[0]
    } else {
        user = 'world'
    }

    data = {
        id: '',
        brand: 'intel',
        name: 'intel intel',
        cpu: {
            brand: 'intel',
            name: 'intel',
            number_cores: 8,
            number_threads: 8,
            min_ghz: 500,
            max_ghz: 500,
        },
        ram: {
            value: 500,
            unit: 'GIGABYTE',
        },
        gpus: [
            {
                brand: 'nvidia',
                name: 'nvidia',
                min_ghz: 500,
                max_ghz: 500,
                memory: {
                    value: 500,
                    unit: 'GIGABYTE',
                },
            },
        ],
        storages: [
            {
                driver: 'SSD',
                memory: {
                    value: 10,
                    unit: 'TERABYTE',
                },
            },
        ],
        screen: {
            size_inch: 65,
            resolution: {
                width: 6076,
                height: 3418,
            },
            panel: 'OLED',
            multitouch: true,
        },
        keyboard: {
            layout: 'QWERTY',
            backlit: true,
        },
        weight_kg: 10,
        price_usd: 9999,
        release_year: 2022,
        // updated_at: Date.now()
    }

    await client.CreateLaptop({ laptop: data }, function (err, response) {
        console.log('Laptop:', response)
        client.Find({ id: response.id }, function (err, resp) {
            console.log('Laptop: ', resp.laptop)
        })
    })

    client.Hello({ name: 'Lino' }, function (err, response) {
        console.log('Greeting:', response.reply)
    })
}

main()
