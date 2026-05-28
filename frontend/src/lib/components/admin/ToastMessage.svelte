<script>
  import { toast } from "$lib/stores/toastStore";
</script>

{#if $toast}
  <div role="status" aria-live="polite" class="fixed top-5 right-5 z-[9999] flex items-center gap-2.5 px-5 py-3 rounded-xl text-sm font-medium shadow-xl animate-slideIn
  {$toast.type === 'error'
    ? 'bg-red-500/90 text-white'
    : 'bg-emerald-500/90 text-white'
  }
  "><i class="fa-solid {$toast.type === 'error' ? 'fa-circle-xmark' : 'fa-circle-check'}"></i>
    {$toast.msg}
  </div>
{/if}

<style>
  @keyframes slideIn {
    from { opacity: 0; transform: translateX(1rem); }
    to { opacity: 1; transform: translateX(0); }
  }

  .animate-slideIn { animation: slideIn 0.25s ease-out; }
</style>