package bluetooth

import (
	"fmt"
	"log"
	"time"

	"github.com/paypal/gatt"
	"github.com/paypal/gatt/examples/option"
)

// StartBluetoothScanner inicia o scanner Bluetooth
func StartBluetoothScanner() {
	// Configuração básica do dispositivo Bluetooth
	d, err := gatt.NewDevice(option.DefaultClientOptions...)
	if err != nil {
		log.Fatalf("Erro ao criar dispositivo Bluetooth: %v", err)
		return
	}

	// Manipulador para eventos de descoberta de dispositivo
	d.Handle(gatt.PeripheralDiscovered(func(p gatt.Peripheral, a *gatt.Advertisement, rssi int) {
		fmt.Printf("Dispositivo encontrado: %s, RSSI: %d\n", p.ID(), rssi)
		// Verificar se há vulnerabilidades
		if hasVulnerabilities(a) {
			fmt.Println("!!! Dispositivo vulnerável encontrado !!!")
			fmt.Printf("ID: %s, Nome: %s\n", p.ID(), a.LocalName)
			// Aqui você pode adicionar lógica para lidar com dispositivos vulneráveis
			// Por exemplo, notificar administradores ou tomar medidas de segurança
		}
	}))

	// Iniciar dispositivo Bluetooth
	d.Init(onStateChanged)

	// Manter o programa ativo indefinidamente para continuar escaneando
	select {}
}

// Função de callback para estado inicializado do dispositivo Bluetooth
func onStateChanged(d gatt.Device, s gatt.State) {
	switch s {
	case gatt.StatePoweredOn:
		fmt.Println("Dispositivo Bluetooth está pronto. Iniciando escaneamento...")
		// Iniciar escaneamento de dispositivos
		d.Scan([]gatt.UUID{}, true)
	default:
		fmt.Printf("Estado do dispositivo Bluetooth: %s\n", s)
	}
}

// Função para verificar se há vulnerabilidades com base nos serviços e características do Advertisement
func hasVulnerabilities(a *gatt.Advertisement) bool {
	// Exemplo simples: verificar por um serviço ou característica específica conhecida por ser vulnerável
	for _, service := range a.Services {
		if service.UUID.Equal(gatt.MustParseUUID("0000180d-0000-1000-8000-00805f9b34fb")) {
			// Exemplo: serviço Heart Rate, conhecido por ser comum e não necessariamente vulnerável, apenas um exemplo
			return true
		}
	}

	return false
}
