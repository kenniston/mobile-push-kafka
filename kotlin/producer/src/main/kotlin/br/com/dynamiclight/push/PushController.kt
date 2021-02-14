package br.com.dynamiclight.push

import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("v1")
class PushController(val pushService: PushService) {

    @PostMapping("/push/send")
    private fun send(@RequestBody message: Push) : ResponseEntity<*> {
        pushService.send(message)
        return ResponseEntity.accepted().build<ResponseEntity<*>>()
    }

}