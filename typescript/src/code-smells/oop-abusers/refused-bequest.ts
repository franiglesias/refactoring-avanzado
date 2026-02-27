class BaseController {
  start(): void {
    console.log('starting')
  }

  stop(): void {
    console.log('stopping')
  }

  reset(): void {
    console.log('resetting')
  }
}

export class ReadOnlyController extends BaseController {
  start(): void {
  }

  stop(): void {
  }
}

export function demoRefusedBequest(readonly: boolean): void {
  const controller: BaseController = readonly ? new ReadOnlyController() : new BaseController()
  controller.start()
  controller.stop()
}
